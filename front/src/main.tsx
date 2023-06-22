import { ChakraProvider } from '@chakra-ui/react';
import React from 'react';
import ReactDOM from 'react-dom/client';

import { RouterProvider, createBrowserRouter } from 'react-router-dom';

import { QueryClient, QueryClientProvider } from 'react-query';

import { ApiView } from './pages/admin-pages';
import { Error404, Login, Profile, Register, UserRoot } from './pages/user-pages';
import { AuthProvider } from './providers/auth-context';

const router = createBrowserRouter([
  {
    path: '/',
    element: <UserRoot />,
    errorElement: <Error404 />,
    children: [
      {
        path: 'register',
        element: <Register />,
      },
      {
        path: 'login',
        element: <Login />,
      },
      {
        path: 'profile',
        element: <Profile />,
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
      <ChakraProvider>
        <QueryClientProvider client={queryClient}>
          <RouterProvider router={router} />
        </QueryClientProvider>
      </ChakraProvider>
    </AuthProvider>
  </React.StrictMode>,
);
