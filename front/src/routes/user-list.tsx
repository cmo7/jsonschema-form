import { useQuery } from 'react-query';
import UserLayout from '../layouts/user/user-layout';
import UserBadge from '../components/user/user-badge';
import { Box, Center, Spinner } from '@chakra-ui/react';

interface UserListElement {
  id: number;
  created_at: string;
  updated_at: string;
  email: string;
  friends: UserListElement[];
}

export default function UserList() {
  const userList = useQuery<UserListElement[]>('userList', async () => {
    const response = await fetch('http://localhost:8080/api/user/all');
    const jsonData = await response.json();
    return jsonData;
  });

  if (userList.isLoading)
    return (
      <UserLayout>
        <Box minH="100vh">
          <Loader />
        </Box>
      </UserLayout>
    );
  if (userList.isError) return <h1>Error</h1>;
  if (!userList.data) return <h1>No data</h1>;
  return (
    <UserLayout>
      {userList.data.map((u) => (
        <UserBadge key={u.id} id={u.id} email={u.email} friends={u.friends} />
      ))}
    </UserLayout>
  );
}

function Loader() {
  return (
    <Center>
      <Spinner thickness="4px" speed="0.65s" emptyColor="gray.200" color="blue.500" size="xl" />
    </Center>
  );
}
