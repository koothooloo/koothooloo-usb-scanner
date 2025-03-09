# Koothooloo - USB Device Scanner

This is a USB device browser application with a Go backend and React frontend. It allows users to view detailed information about USB devices connected to their system.

## Features

- Real-time USB device detection and monitoring
- Detailed device information display including:
  - Manufacturer and product details
  - Technical specifications (vendor ID, product ID, device class)
  - Power specifications
  - Connection details
- Automatic USB ID lookup using online database
- Responsive UI built with Chakra UI

## Architecture

- **Backend**: Go application using [gousb](https://github.com/google/gousb) for USB device access
- **Frontend**: React application with Vite and Chakra UI
- **API**: Simple RESTful API to communicate device information

## Prerequisites

- Go 1.16+
- Node.js 14+ and pnpm
- USB devices to test with

## Installation and Running

### Backend (Go)

1. Clone the repository:

   ```bash
   git clone https://github.com/koothooloo/koothooloo-usb-scanner.git
   cd koothooloo-usb-scanner
   ```

2. Install Go dependencies:

   ```bash
   go mod download
   ```

3. Run the backend server:

   ```bash
   go run main.go usbids.go
   ```

The backend server will start on port 8080 by default.

### Frontend (React)

1. Navigate to the frontend directory:

   ```bash
   cd frontend
   ```

2. Install dependencies:

   ```bash
   pnpm install
   ```

3. Start the development server:

   ```bash
   pnpm dev
   ```

The frontend development server will start on port 5173 by default and can be accessed at <http://localhost:5173>.

## Building for Production

### Frontend

```bash
cd frontend
pnpm build
```

This will create a production build in the `frontend/dist` directory.

### Backend

```bash
go build -o koothooloo main.go usbids.go
```

This will create an executable named `koothooloo` that you can run on your system.

## Usage

1. Start both the backend and frontend servers
2. Open a web browser and navigate to <http://localhost:5173>
3. The application will display all USB devices connected to your system
4. Click on any device to see detailed information
5. The list automatically refreshes every 5 seconds

## Technical Details

- The Go backend uses the USB IDs database to match vendor and product IDs with human-readable names
- The application caches USB ID information locally to minimize network requests
- Device information is fetched through gousb library which provides low-level access to USB devices
- The frontend is built with React and uses Chakra UI components for a clean, responsive interface

## License

This project is open source and available under the MIT License.
