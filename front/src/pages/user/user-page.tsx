import { Avatar, Center, VStack } from '@chakra-ui/react';
import { useQuery } from 'react-query';
import { useParams } from 'react-router-dom';
import { getUserById } from '../../api/user-requests';
import useTitle from '../../hooks/title';

export default function UserPage() {
  useTitle('Perfil de usuario');
  const id = useParams().id;

  const user = useQuery('user', async () => {
    return getUserById(id);
  });

  if (user.isLoading) return <div>Loading...</div>;

  if (user.data === null) return <div>Usuario no encontrado</div>;

  return (
    <Center>
      <VStack>
        <Avatar size="2xl" name={user.data.name} src={user.data.avatar} />
        <h1>{user.data.name}</h1>
        <h2>{user.data.email}</h2>
      </VStack>
    </Center>
  );
}
