import { ReactNode } from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from '../hooks/auth';

interface GuestRouteProps {
  children: ReactNode;
}

export default function GuestRoute({ children }: GuestRouteProps) {
  const auth = useAuth();

  if (auth.isAuthenticated()) {
    return <Navigate to="/" />;
  }

  return <>{children}</>;
}
