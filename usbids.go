package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	usbIDsAPI   = "https://apps.sebastianlang.net/usb-ids"
	cacheFile   = "usb-ids-cache.json"
	cacheExpiry = 24 * time.Hour
)

type USBIDsCache struct {
	Vendors   map[string]string            `json:"vendors"`
	Products  map[string]map[string]string `json:"products"`
	UpdatedAt time.Time                    `json:"updatedAt"`
	mutex     sync.RWMutex                 `json:"-"`
}

var (
	usbIDsCache *USBIDsCache
	cacheOnce   sync.Once
)

func getUSBIDsCache() *USBIDsCache {
	cacheOnce.Do(func() {
		usbIDsCache = &USBIDsCache{
			Vendors:  make(map[string]string),
			Products: make(map[string]map[string]string),
		}
		loadCache()
	})
	return usbIDsCache
}

func loadCache() {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		cacheDir = "."
	}
	cachePath := filepath.Join(cacheDir, "koothooloo", cacheFile)

	// Ensure cache directory exists
	if err := os.MkdirAll(filepath.Dir(cachePath), 0755); err != nil {
		log.Printf("Error creating cache directory: %v", err)
		return
	}

	data, err := os.ReadFile(cachePath)
	if err != nil {
		// If cache doesn't exist or can't be read, fetch fresh data
		updateCache()
		return
	}

	if err := json.Unmarshal(data, usbIDsCache); err != nil {
		log.Printf("Error unmarshaling cache: %v", err)
		updateCache()
		return
	}

	// Check if cache is expired
	if time.Since(usbIDsCache.UpdatedAt) > cacheExpiry {
		updateCache()
	}
}

func updateCache() {
	usbIDsCache.mutex.Lock()
	defer usbIDsCache.mutex.Unlock()

	resp, err := http.Get(usbIDsAPI)
	if err != nil {
		log.Printf("Error fetching USB IDs: %v", err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v", err)
		return
	}

	var apiData struct {
		Vendors  map[string]string            `json:"vendors"`
		Products map[string]map[string]string `json:"products"`
	}

	if err := json.Unmarshal(data, &apiData); err != nil {
		log.Printf("Error parsing USB IDs data: %v", err)
		return
	}

	usbIDsCache.Vendors = apiData.Vendors
	usbIDsCache.Products = apiData.Products
	usbIDsCache.UpdatedAt = time.Now()

	// Save to cache file
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		cacheDir = "."
	}
	cachePath := filepath.Join(cacheDir, "koothooloo", cacheFile)

	cacheData, err := json.Marshal(usbIDsCache)
	if err != nil {
		log.Printf("Error marshaling cache: %v", err)
		return
	}

	if err := os.WriteFile(cachePath, cacheData, 0644); err != nil {
		log.Printf("Error writing cache: %v", err)
	}
}

func lookupVendor(vendorID string) string {
	cache := getUSBIDsCache()
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	if name, ok := cache.Vendors[vendorID]; ok {
		return name
	}
	return ""
}

func lookupProduct(vendorID, productID string) string {
	cache := getUSBIDsCache()
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	if products, ok := cache.Products[vendorID]; ok {
		if name, ok := products[productID]; ok {
			return name
		}
	}
	return ""
}
