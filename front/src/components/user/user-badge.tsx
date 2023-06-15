import { PropsWithChildren } from 'react';

import { Avatar, Box, Center, HStack, ListItem, UnorderedList } from '@chakra-ui/react';
import { Link } from 'react-router-dom';

export interface UserBadgeProps {
  id: number;
  email: string;
  friends: UserBadgeProps[];
}

export default function UserBadge(props: PropsWithChildren<UserBadgeProps>) {
  const { id, email, friends } = props;
  return (
    <Center py={6}>
      <Box maxW="sm" borderWidth="1px" borderRadius="lg" overflow="hidden">
        <Link to={`/user/${id}`}>
          <HStack>
            <Avatar name={email} />
            <h4>{email}</h4>
          </HStack>
        </Link>
        <UnorderedList>
          {friends.map((f) => (
            <FriendBadge key={f.id} id={f.id} email={f.email} friends={f.friends} />
          ))}
        </UnorderedList>
      </Box>
    </Center>
  );
}

function FriendBadge(props: PropsWithChildren<UserBadgeProps>) {
  const { id, email } = props;
  return (
    <ListItem>
      <UserLink id={id} email={email} />
    </ListItem>
  );
}

function UserLink(props: { id: number; email: string }) {
  return <Link to={`/user/${props.id}`}>{props.email}</Link>;
}
