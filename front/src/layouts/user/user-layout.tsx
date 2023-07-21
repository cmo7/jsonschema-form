import { Container, Spacer, VStack } from '@chakra-ui/react';
import { PropsWithChildren } from 'react';
import Footer from '../../components/footer';
import Header from '../../components/header';

export default function UserLayout({ children }: PropsWithChildren) {
  return (
    <VStack minH={'100vh'} align={'stretch'} justify={'space-between'}>
      <Header />
      <Container maxW={'container.xl'}>{children}</Container>
      <Spacer />
      <Footer />
    </VStack>
  );
}
