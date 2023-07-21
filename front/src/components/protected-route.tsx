import { ReactNode } from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from '../hooks/auth';

interface ProtectedRouteProps {
  children: ReactNode;
}

export default function ProtectedRoute({ children }: ProtectedRouteProps) {
  const auth = useAuth();

  if (!auth.isAuthenticated()) {
    return <Navigate to="/login" />;
  }

  return <>{children}</>;
}
