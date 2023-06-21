import { Box, Flex, VStack } from '@chakra-ui/react';
import { PropsWithChildren } from 'react';
import Footer from '../../components/footer';
import Header from '../../components/header';
import Sidebar from '../../components/sidebar';

export default function UserLayout(props: PropsWithChildren) {
  return (
    <>
      <Flex>
        <Sidebar />
        <Box flex="1">
          <VStack align="stretch">
            <Header />
            <Box>{props.children}</Box>
            <Footer />
          </VStack>
        </Box>
      </Flex>
    </>
  );
}
