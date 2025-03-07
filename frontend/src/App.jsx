import { Box, Container, Heading, SimpleGrid, Text, VStack, useToast } from '@chakra-ui/react';
import { useEffect, useState } from 'react';
import { FaUsb } from 'react-icons/fa';

function App() {
  const [devices, setDevices] = useState([]);
  const toast = useToast();

  const fetchDevices = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/devices');
      const data = await response.json();
      setDevices(data);
    } catch (error) {
      toast({
        title: 'Error fetching devices',
        description: error.message,
        status: 'error',
        duration: 5000,
        isClosable: true,
      });
    }
  };

  useEffect(() => {
    fetchDevices();
    const interval = setInterval(fetchDevices, 5000);
    return () => clearInterval(interval);
  }, []);

  return (
    <Box minH="100vh" bg="gray.50" py={8}>
      <Container maxW="container.xl">
        <VStack spacing={8} align="stretch">
          <Box textAlign="center">
            <Heading size="xl" mb={2} display="flex" alignItems="center" justifyContent="center">
              <FaUsb style={{ marginRight: '12px' }} />
              USB Device Browser
            </Heading>
            <Text color="gray.600">Connected USB Devices: {devices.length}</Text>
          </Box>

          <SimpleGrid columns={{ base: 1, md: 2, lg: 3 }} spacing={6}>
            {devices.map((device, index) => (
              <Box
                key={index}
                bg="white"
                p={6}
                rounded="lg"
                shadow="md"
                borderWidth="1px"
                _hover={{ transform: 'translateY(-2px)', shadow: 'lg' }}
                transition="all 0.2s"
              >
                <VStack align="stretch" spacing={3}>
                  <Heading size="md" color="blue.600">
                    Device {index + 1}
                  </Heading>
                  <Box>
                    <Text fontWeight="bold">Manufacturer:</Text>
                    <Text>{device.manufacturer || 'N/A'}</Text>
                  </Box>
                  <Box>
                    <Text fontWeight="bold">Product:</Text>
                    <Text>{device.product || 'N/A'}</Text>
                  </Box>
                  <Box>
                    <Text fontWeight="bold">Vendor ID:</Text>
                    <Text>{device.vendorId}</Text>
                  </Box>
                  <Box>
                    <Text fontWeight="bold">Product ID:</Text>
                    <Text>{device.productId}</Text>
                  </Box>
                  <Box>
                    <Text fontWeight="bold">Serial Number:</Text>
                    <Text>{device.serial || 'N/A'}</Text>
                  </Box>
                  <Box>
                    <Text fontWeight="bold">Device Class:</Text>
                    <Text>{device.deviceClass || 'N/A'}</Text>
                  </Box>
                  <Box>
                    <Text fontWeight="bold">Device SubClass:</Text>
                    <Text>{device.deviceSubClass || 'N/A'}</Text>
                  </Box>
                  <Box>
                    <Text fontWeight="bold">Device Protocol:</Text>
                    <Text>{device.deviceProtocol || 'N/A'}</Text>
                  </Box>
                  <Box>
                    <Text fontWeight="bold">Speed:</Text>
                    <Text>{device.speed || 'N/A'}</Text>
                  </Box>
                  <Box>
                    <Text fontWeight="bold">Location ID:</Text>
                    <Text>{device.locationId || 'N/A'}</Text>
                  </Box>
                  <Box>
                    <Text fontWeight="bold">Current Available:</Text>
                    <Text>{device.currentAvailable ? `${device.currentAvailable}mA` : 'N/A'}</Text>
                  </Box>
                  <Box>
                    <Text fontWeight="bold">Current Required:</Text>
                    <Text>{device.currentRequired ? `${device.currentRequired}mA` : 'N/A'}</Text>
                  </Box>
                  <Box>
                    <Text fontWeight="bold">Extra Operating Current:</Text>
                    <Text>{device.extraOperatingCurrent ? `${device.extraOperatingCurrent}mA` : 'N/A'}</Text>
                  </Box>
                  <Box>
                    <Text fontWeight="bold">Bus Number:</Text>
                    <Text>{device.busNumber || 'N/A'}</Text>
                  </Box>
                  <Box>
                    <Text fontWeight="bold">Port Numbers:</Text>
                    <Text>{device.portNumbers ? device.portNumbers.join(', ') : 'N/A'}</Text>
                  </Box>
                  <Box>
                    <Text fontWeight="bold">Description:</Text>
                    <Text>{device.description || 'N/A'}</Text>
                  </Box>
                </VStack>
              </Box>
            ))}
          </SimpleGrid>

          {devices.length === 0 && (
            <Box textAlign="center" p={8} bg="white" rounded="lg" shadow="md">
              <Text fontSize="lg" color="gray.600">
                No USB devices found. Connect a device to see it here.
              </Text>
            </Box>
          )}
        </VStack>
      </Container>
    </Box>
  );
}

export default App;