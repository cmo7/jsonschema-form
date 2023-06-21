import { ReactNode, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../hooks/auth';

interface ProtectedRouteProps {
  children: ReactNode;
}

export default function ProtectedRoute({ children }: ProtectedRouteProps) {
  const navigate = useNavigate();
  const auth = useAuth();
  useEffect(() => {
    if (!auth.isAuthenticated()) {
      navigate('/login');
    }
  }, [auth, navigate]);
  if (!auth.isAuthenticated()) {
    return null;
  }

  return <>{children}</>;
}
