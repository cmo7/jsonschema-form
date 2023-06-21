import { ReactNode, createContext, useMemo } from 'react';
import useLocalStorage from '../hooks/localstorage';
import { UserResponse } from '../types/generated/models';

export interface AuthData {
  userData: UserResponse;
  token: string;
}

export type AuthContextType = {
  auth: AuthData | null;
  user: () => UserResponse;
  token: () => string;
  login: (token: string, user: UserResponse) => void;
  logout: () => void;
  isAuthenticated: () => boolean;
};

type AuthProviderProps = {
  children: ReactNode;
};

export const AuthContext = createContext<AuthContextType>(null);

export function AuthProvider({ children }: AuthProviderProps) {
  const [auth, setAuth] = useLocalStorage<AuthData>('auth', null);

  const value = useMemo(() => {
    return {
      auth,
      login: (token: string, userData: UserResponse) => {
        setAuth({
          userData: userData ? userData : null,
          token,
        });
      },
      logout: () => {
        setAuth(null);
      },
      isAuthenticated: () => {
        return auth !== null;
      },
      user: () => {
        return auth ? auth.userData : null;
      },
      token: () => {
        return auth ? auth.token : null;
      },
    };
  }, [auth, setAuth]);

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}
