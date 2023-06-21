import { Avatar } from '@chakra-ui/react';
import { useQuery } from 'react-query';
import { getCurrentUser } from '../api/auth-requests';
import { useAuth } from '../hooks/auth';
import { Loading } from './loading';

export default function Profile() {
  const auth = useAuth();
  const token = auth.authData.token;

  const userData = useQuery('profile', () => {
    return getCurrentUser(token);
  });

  if (userData.isLoading) return <Loading />;
  if (userData.isError) return <div>error</div>;
  if (!userData.data) return <div>no data</div>;

  const user = userData.data;

  console.log(user, 'user');

  return (
    <div>
      <h1>Profile</h1>
      <Avatar size="2xl" name={user.name} src={user.avatar} />
      <p>Username: {user.name}</p>
    </div>
  );
}
