import { Box, Flex, VStack } from '@chakra-ui/react';
import { PropsWithChildren } from 'react';
import Footer from '../../components/footer';
import Header from '../../components/header';

export default function UserLayout(props: PropsWithChildren) {
  return (
    <>
      <Flex>
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
