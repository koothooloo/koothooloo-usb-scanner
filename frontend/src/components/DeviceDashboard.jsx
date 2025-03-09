import { Modal, ModalOverlay, ModalContent, ModalHeader, ModalBody, ModalCloseButton, Box, Grid, GridItem, Text, VStack, Heading, Divider, Badge } from '@chakra-ui/react';

function DeviceDashboard({ isOpen, onClose, device }) {
  if (!device) return null;

  const sections = [
    {
      title: 'Device Information',
      items: [
        { label: 'Manufacturer', value: device.manufacturer },
        { label: 'Product', value: device.product },
        { label: 'Serial Number', value: device.serial },
        { label: 'Description', value: device.description }
      ]
    },
    {
      title: 'Technical Details',
      items: [
        { label: 'Vendor ID', value: device.vendorId },
        { label: 'Product ID', value: device.productId },
        { label: 'Device Class', value: device.deviceClass },
        { label: 'Device SubClass', value: device.deviceSubClass },
        { label: 'Device Protocol', value: device.deviceProtocol }
      ]
    },
    {
      title: 'Power Specifications',
      items: [
        { label: 'Current Available', value: device.currentAvailable ? `${device.currentAvailable}mA` : 'N/A' },
        { label: 'Current Required', value: device.currentRequired ? `${device.currentRequired}mA` : 'N/A' },
        { label: 'Extra Operating Current', value: device.extraOperatingCurrent ? `${device.extraOperatingCurrent}mA` : 'N/A' },
        { label: 'Max Power', value: device.maxPower }
      ]
    },
    {
      title: 'Connection Details',
      items: [
        { label: 'Speed', value: device.speed },
        { label: 'Bus Number', value: device.busNumber },
        { label: 'Port Numbers', value: device.portNumbers ? device.portNumbers.join(', ') : 'N/A' },
        { label: 'Location ID', value: device.locationId }
      ]
    }
  ];

  return (
    <Modal isOpen={isOpen} onClose={onClose} size="xl">
      <ModalOverlay />
      <ModalContent maxW="900px">
        <ModalHeader>
          <Heading size="lg">{device.product || 'USB Device Details'}</Heading>
          <Badge colorScheme="blue" mt={2}>{device.deviceClass}</Badge>
        </ModalHeader>
        <ModalCloseButton />
        <ModalBody pb={6}>
          <Grid templateColumns={{ base: '1fr', md: 'repeat(2, 1fr)' }} gap={6}>
            {sections.map((section, idx) => (
              <GridItem key={idx}>
                <VStack align="stretch" spacing={4}>
                  <Heading size="md" color="blue.600">
                    {section.title}
                  </Heading>
                  <Divider />
                  {section.items.map((item, itemIdx) => (
                    <Box key={itemIdx}>
                      <Text fontWeight="bold" color="gray.600">
                        {item.label}
                      </Text>
                      <Text>{item.value || 'N/A'}</Text>
                    </Box>
                  ))}
                </VStack>
              </GridItem>
            ))}
          </Grid>
        </ModalBody>
      </ModalContent>
    </Modal>
  );
}

export default DeviceDashboard;