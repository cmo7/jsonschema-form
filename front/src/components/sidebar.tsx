import { Box, Button, ListItem, UnorderedList, useColorModeValue } from '@chakra-ui/react';
import { useMutation } from 'react-query';
import { Link, useNavigate } from 'react-router-dom';
import { logout } from '../api/auth-requests';
import { useAuth } from '../hooks/auth';

export default function Sidebar() {
  return (
    <Box minW="250px" minH="100vh" bg={useColorModeValue('gray.100', 'gray.900')}>
      <UnorderedList>
        <HomeLink />
        <AuthSection />
      </UnorderedList>
    </Box>
  );
}

function AuthSection() {
  const auth = useAuth();
  if (!auth.isAuthenticated()) {
    return (
      <>
        <LoginLink />
        <RegisterLink />
      </>
    );
  } else {
    return (
      <>
        <ProfileLink />
        <LogOutLink />
      </>
    );
  }
}

function HomeLink() {
  return (
    <ListItem>
      <Link to="/">Home</Link>
    </ListItem>
  );
}

function LogOutLink() {
  const auth = useAuth();
  const navigate = useNavigate();
  const mutation = useMutation({
    mutationFn: async () => {
      logout();
      auth.logout();
    },
    onSuccess: () => {
      console.log('logged out');
      navigate('/');
    },
  });
  return (
    <ListItem>
      <Button
        variant={'link'}
        onClick={() => {
          mutation.mutate();
        }}
      >
        Logout
      </Button>
    </ListItem>
  );
}

function LoginLink() {
  return (
    <ListItem>
      <Link to="/login">Login</Link>
    </ListItem>
  );
}

function RegisterLink() {
  return (
    <ListItem>
      <Link to="/register">Register</Link>
    </ListItem>
  );
}

function ProfileLink() {
  return (
    <ListItem>
      <Link to="/profile">Profile</Link>
    </ListItem>
  );
}
