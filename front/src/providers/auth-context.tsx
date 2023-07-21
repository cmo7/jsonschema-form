import { ReactNode, createContext, useMemo } from 'react';
import useLocalStorage from '../hooks/localstorage';
import { UserDTO } from '../types/generated/models';

export interface AuthData {
  userData: UserDTO;
  token: string;
}

export type AuthContextType = {
  /**
   * The auth object stores the user data and the token.
   */
  auth: AuthData | null;
  /**
   * The user() function returns the user data from the auth object.
   * @returns UserDTO
   */
  user: () => UserDTO;
  /**
   * The token() function returns the token from the auth object.
   * @returns string
   */
  token: () => string;
  /**
   * The login() function receives a token and a user object and stores them in the auth object.
   * @param token
   * @param user
   * @returns
   */
  login: (token: string, user: UserDTO) => void;
  /**
   * The logout() function removes the auth object from the local storage.
   * @returns
   */
  logout: () => void;
  /**
   * The isAuthenticated() function checks if the auth object is null or not.
   * @returns boolean
   */
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
      login: (token: string, userData: UserDTO) => {
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
