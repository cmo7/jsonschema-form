import { Avatar, Box } from '@chakra-ui/react';
import { useAuth } from '../hooks/auth';

export default function Header() {
  const auth = useAuth();
  return (
    <Box bg="coral">
      <div className="header">Header</div>
      <Avatar />
    </Box>
  );
}
