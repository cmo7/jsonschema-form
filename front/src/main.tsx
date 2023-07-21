import { ChakraProvider } from '@chakra-ui/react';
import React from 'react';
import ReactDOM from 'react-dom/client';

import { RouterProvider, createBrowserRouter } from 'react-router-dom';

import { QueryClient, QueryClientProvider } from 'react-query';

import GuestRoute from './components/guest-route';
import ProtectedRoute from './components/protected-route';
import { ApiView } from './pages/admin-pages';
import { Login, Profile, Register, UserList, UserPage, UserRoot } from './pages/user-pages';
import { Error404 } from './pages/user/error404';
import { AuthProvider } from './providers/auth-context';
import theme from './theme';

const router = createBrowserRouter([
  {
    path: '/',
    element: <UserRoot />,
    errorElement: <Error404 />,
    children: [
      {
        path: 'register',
        element: (
          <GuestRoute>
            <Register />
          </GuestRoute>
        ),
      },
      {
        path: 'login',
        element: (
          <GuestRoute>
            <Login />
          </GuestRoute>
        ),
      },
      {
        path: 'profile',
        element: (
          <ProtectedRoute>
            <Profile />
          </ProtectedRoute>
        ),
      },
      {
        path: 'users',
        element: <UserList />,
      },
      {
        path: 'user/:id',
        element: <UserPage />,
      },
    ],
  },
  {
    path: '/admin',
    element: <UserRoot />,
    errorElement: <Error404 />,
    children: [
      {
        path: 'api-view',
        element: <ApiView />,
      },
    ],
  },
]);

const queryClient = new QueryClient();

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <AuthProvider>
      <ChakraProvider theme={theme}>
        <QueryClientProvider client={queryClient}>
          <RouterProvider router={router} />
        </QueryClientProvider>
      </ChakraProvider>
    </AuthProvider>
  </React.StrictMode>,
);
