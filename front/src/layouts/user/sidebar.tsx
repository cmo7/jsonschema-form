import { Box, useColorModeValue } from "@chakra-ui/react";

export default function Sidebar() {
    return (
        <Box minW='250px' minH="100vh" bg={useColorModeValue('gray.100', 'gray.900')}>
            <p>Sidebar</p>
        </Box>
    )
}