import { useParams } from 'react-router-dom';
import UserLayout from '../layouts/user/user-layout';
import { useMutation, useQuery } from 'react-query';
import { Button, Center, Skeleton, SkeletonCircle } from '@chakra-ui/react';
import UserBadge from '../components/user/user-badge';

export default function UserProfile() {
  const { userId } = useParams<'userId'>();

  const user = useQuery('user', async () => {
    const response = await fetch(`http://localhost:8080/api/user/${userId}`);
    const jsonData = await response.json();
    return jsonData;
  });

  const deleteUser = useMutation({
    mutationFn: async () => {
      const response = await fetch(`http://localhost:8080/api/user/${userId}`, {
        method: 'DELETE',
      });
      const jsonData = await response.json();
      return jsonData;
    },
    onSuccess: () => {
      window.location.href = '/users';
    },
  });

  if (user.isLoading)
    return (
      <UserLayout>
        <SkeletonCircle size="10" />
        <Skeleton height="100px" />
      </UserLayout>
    );

  if (user.isError)
    return (
      <UserLayout>
        <h1>Error</h1>
      </UserLayout>
    );
  if (!user.data)
    return (
      <UserLayout>
        <h1>No data</h1>
      </UserLayout>
    );

  const { id, email, friends } = user.data;
  return (
    <UserLayout>
      <UserBadge id={id} email={email} friends={friends} />
      <Center>
        <Button
          variant="solid"
          colorScheme="red"
          onClick={() => {
            deleteUser.mutate();
          }}
        >
          Delete
        </Button>
      </Center>
    </UserLayout>
  );
}
