import { Outlet } from 'react-router-dom';
import UserLayout from '../../layouts/user/user-layout';

export default function UserRoot() {
  return (
    <UserLayout>
      <Outlet />
    </UserLayout>
  );
}
