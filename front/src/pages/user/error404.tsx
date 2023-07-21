import { Center } from '@chakra-ui/react';
import useTitle from '../../hooks/title';

export function Error404() {
  useTitle('404');
  return <Center>404</Center>;
}
