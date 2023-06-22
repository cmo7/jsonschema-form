import { Avatar } from '@chakra-ui/react';
import ProtectedRoute from '../../components/protected-route';
import { useAuth } from '../../hooks/auth';

export default function Profile() {
  const user = useAuth().user();

  console.log(user, 'user');

  return (
    <ProtectedRoute>
      <div>
        <h1>Profile</h1>
        <Avatar size="2xl" name={user.name} src={user.avatar} />
        <p>Username: {user.name}</p>
      </div>
    </ProtectedRoute>
  );
}
