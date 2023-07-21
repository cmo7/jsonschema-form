import { Center } from '@chakra-ui/react';
import useTitle from '../../hooks/title';

export function ServerError() {
  useTitle('Error en el servidor');
  return <Center>Error en el servidor</Center>;
}
