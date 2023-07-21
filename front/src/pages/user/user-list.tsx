import { Avatar, HStack, Table, TableContainer, Tbody, Td, Text, Thead, Tr, VStack } from '@chakra-ui/react';
import React, { useState } from 'react';
import { useQuery } from 'react-query';
import { Link } from 'react-router-dom';
import { endpoints } from '../../api/endpoints';
import { http } from '../../api/http-methods';
import Pagination from '../../components/pagination';
import useTitle from '../../hooks/title';
import { ApiResponse } from '../../types/api-response';
import { Page } from '../../types/generated/controllers';
import { UserDTO } from '../../types/generated/models';
import { Loading } from './loading';
import { NoData } from './no-data';

export const UserList: React.FC = () => {
  useTitle('Usuarios');
  const [page, setPage] = useState(1);
  const [pageSize, setPageSize] = useState(10);

  const users = useQuery<ApiResponse<Page>>({
    queryKey: ['users', page],
    queryFn: () => {
      const queryString = `page=${page}&size=${pageSize}`;
      return http.get<ApiResponse<Page>>(endpoints.user + '?' + queryString, {});
    },
    keepPreviousData: true,
  });

  if (users.isLoading) {
    return <Loading />;
  }

  if (users.data == null) {
    return <NoData />;
  }

  console.log(users.data.data);
  return (
    <>
      <TableContainer>
        <Table variant="simple">
          <Thead>
            <Td>Nombre</Td>
          </Thead>
          <Tbody>
            {users.data.data.content.map((u) => {
              return <UserRow key={u.id} user={u} />;
            })}
          </Tbody>
        </Table>
      </TableContainer>
      <Pagination
        page={page}
        pageSize={pageSize}
        setPage={setPage}
        setPageSize={setPageSize}
        totalPages={users.data.data.total / pageSize}
      />
    </>
  );
};

const UserRow: React.FC<{ user: UserDTO }> = ({ user }) => {
  return (
    <Tr>
      <Td>
        <Link to={'/user/' + user.id}>
          <VStack>
            <HStack spacing={3}>
              <Avatar src={user.avatar} />
              <Text>{user.name}</Text>
            </HStack>
          </VStack>
        </Link>
      </Td>
    </Tr>
  );
};
