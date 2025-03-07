package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/gousb"
)

type USBDevice struct {
	VendorID         string `json:"vendorId"`
	ProductID        string `json:"productId"`
	Serial           string `json:"serial"`
	Desc             string `json:"description"`
	Manufacturer     string `json:"manufacturer"`
	Product          string `json:"product"`
	BusNumber        int    `json:"busNumber"`
	PortNumbers      []int  `json:"portNumbers"`
	DeviceClass      string `json:"deviceClass"`
	DeviceSubClass   string `json:"deviceSubClass"`
	DeviceProtocol   string `json:"deviceProtocol"`
	Speed            string `json:"speed"`
	MaxPower         string `json:"maxPower"`
	LocationID       string `json:"locationId"`
	CurrentAvailable int    `json:"currentAvailable"`
	CurrentRequired  int    `json:"currentRequired"`
	ExtraOpCurrent   int    `json:"extraOperatingCurrent"`
}

func getUSBDevices() []USBDevice {
	ctx := gousb.NewContext()
	defer ctx.Close()

	devices, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		return true
	})

	if err != nil {
		log.Printf("Error getting USB devices: %v", err)
		return []USBDevice{}
	}

	var usbDevices []USBDevice
	for _, dev := range devices {
		// Get device configurations
		config, err := dev.Config(1)
		if err != nil {
			log.Printf("Error getting device config: %v", err)
		}

		serial, err := dev.SerialNumber()
		if err != nil {
			log.Printf("Error getting serial number: %v", err)
			serial = "unknown"
		}

		manufacturer, err := dev.Manufacturer()
		if err != nil {
			log.Printf("Error getting manufacturer: %v", err)
			manufacturer = "unknown"
		}

		product, err := dev.Product()
		if err != nil {
			log.Printf("Error getting product: %v", err)
			product = "unknown"
		}

		deviceClass := getDeviceClass(dev.Desc.Class)
		deviceSubClass := fmt.Sprintf("0x%02x", dev.Desc.SubClass)
		deviceProtocol := fmt.Sprintf("0x%02x", dev.Desc.Protocol)
		speed := getSpeedString(dev.Desc.Speed)
		maxPower := "Unknown"
		locationID := fmt.Sprintf("0x%08x", dev.Desc.Bus)

		// Calculate current values based on USB specifications and device speed
		currentAvailable := 500 // Default USB 2.0 available current in mA
		switch dev.Desc.Speed {
		case gousb.SpeedSuper:
			currentAvailable = 900 // USB 3.0 provides up to 900mA
		case gousb.SpeedHigh:
			currentAvailable = 500 // USB 2.0 provides up to 500mA
		case gousb.SpeedFull, gousb.SpeedLow:
			currentAvailable = 100 // USB 1.x provides up to 100mA
		}

		currentRequired := 0 // Default to 0 since we can't get MaxPower directly
		extraOpCurrent := 0

		// Get additional configuration details if available
		if config != nil {
			// Get configuration descriptor
			desc := config.Desc
			maxPower = fmt.Sprintf("%dmA", desc.MaxPower)
			currentRequired = int(desc.MaxPower)
			defer config.Close()
		}

		vendorID := dev.Desc.Vendor.String()
		productID := dev.Desc.Product.String()

		// Convert single port number to slice
		portNumbers := []int{dev.Desc.Port}

		device := USBDevice{
			VendorID:         vendorID,
			ProductID:        productID,
			Serial:           serial,
			Desc:             dev.String(),
			Manufacturer:     manufacturer,
			Product:          product,
			BusNumber:        dev.Desc.Bus,
			PortNumbers:      portNumbers,
			DeviceClass:      deviceClass,
			DeviceSubClass:   deviceSubClass,
			DeviceProtocol:   deviceProtocol,
			Speed:            speed,
			MaxPower:         maxPower,
			LocationID:       locationID,
			CurrentAvailable: currentAvailable,
			CurrentRequired:  currentRequired,
			ExtraOpCurrent:   extraOpCurrent,
		}
		usbDevices = append(usbDevices, device)
		dev.Close()
	}

	return usbDevices
}

func getDeviceClass(class gousb.Class) string {
	switch class {
	case gousb.ClassAudio:
		return "Audio"
	case gousb.ClassComm:
		return "Communications"
	case gousb.ClassHID:
		return "HID"
	case gousb.ClassPrinter:
		return "Printer"
	case gousb.ClassHub:
		return "Hub"
	case gousb.ClassData:
		return "Data"
	case gousb.ClassMassStorage:
		return "Mass Storage"
	case gousb.ClassVendorSpec:
		return "Vendor Specific"
	default:
		return fmt.Sprintf("0x%02x", class)
	}
}

func getSpeedString(speed gousb.Speed) string {
	switch speed {
	case gousb.SpeedLow:
		return "Low Speed (1.5 Mbps)"
	case gousb.SpeedFull:
		return "Full Speed (12 Mbps)"
	case gousb.SpeedHigh:
		return "High Speed (480 Mbps)"
	case gousb.SpeedSuper:
		return "Super Speed (5 Gbps)"
	default:
		return "Unknown"
	}
}

func main() {
	r := gin.Default()

	// Enable CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// API endpoints
	r.GET("/api/devices", func(c *gin.Context) {
		devices := getUSBDevices()
		c.JSON(http.StatusOK, devices)
	})

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
