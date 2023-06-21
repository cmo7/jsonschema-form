import { Box, Center } from '@chakra-ui/react';
import { Outlet } from 'react-router-dom';
import UserLayout from '../layouts/user/user-layout';

export default function UserRoot() {
  return (
    <UserLayout>
      <Box minH="100vh">
        <Center>
          <Outlet />
        </Center>
      </Box>
    </UserLayout>
  );
}
