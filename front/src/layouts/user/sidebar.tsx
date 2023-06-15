import { Box, useColorModeValue } from '@chakra-ui/react';
import { Link } from 'react-router-dom';

export default function Sidebar() {
  return (
    <Box minW="250px" minH="100vh" bg={useColorModeValue('gray.100', 'gray.900')}>
      <ul>
        <li>
          <Link to="/">Portada</Link>
        </li>
        <li>
          <Link to="/users">Todos</Link>
        </li>
      </ul>
    </Box>
  );
}
