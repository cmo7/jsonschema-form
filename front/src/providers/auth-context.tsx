import { ReactNode, createContext } from 'react';
import useLocalStorage from '../hooks/localstorage';
import { UserResponse } from '../types/generated/models';

export interface AuthData {
  userData: UserResponse;
  token: string;
}

export type AuthContextType = {
  auth: AuthData | null;
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

  function login(token: string, userData: UserResponse) {
    setAuth({
      userData: userData ? userData : null,
      token,
    });
  }
  function logout() {
    setAuth(null);
  }
  function isAuthenticated() {
    return auth !== null;
  }

  return <AuthContext.Provider value={{ auth, login, logout, isAuthenticated }}>{children}</AuthContext.Provider>;
}
