import { Box } from '@chakra-ui/react';
import CallToActionWithAnnotation from '../components/user/hero';
import UserLayout from '../layouts/user/user-layout';

export default function Root() {
  return (
    <UserLayout>
      <Box>
        <CallToActionWithAnnotation />
      </Box>
    </UserLayout>
  );
}
