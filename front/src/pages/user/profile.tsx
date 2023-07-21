import { Avatar } from '@chakra-ui/react';
import { useAuth } from '../../hooks/auth';
import useTitle from '../../hooks/title';

export default function Profile() {
  const user = useAuth().user();

  useTitle(user.name);
  console.log(user, 'user');

  return (
    <div>
      <h1>Profile</h1>
      <Avatar size="2xl" name={user.name} src={user.avatar} />
      <p>Username: {user.name}</p>
    </div>
  );
}
