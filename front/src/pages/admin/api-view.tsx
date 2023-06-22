import { Table, TableCaption, TableContainer, Td, Th, Thead, Tr } from '@chakra-ui/react';
import { useQuery } from 'react-query';
import { getApiRoutes } from '../../api/admin-requests';
import ProtectedRoute from '../../components/protected-route';
import { useAuth } from '../../hooks/auth';
import { Route } from '../../types/generated/models';

export default function ApiView() {
  const token = useAuth().token();

  const query = useQuery('api-routes', async () => {
    const routes = await getApiRoutes(token);
    if (!routes) {
      return [];
    }
    const routesWithKey = routes.map((route, i) => {
      return {
        ...route,
        key: i,
      };
    });
    return routesWithKey;
  });
  if (query.isLoading) {
    return <div>Loading...</div>;
  }
  if (query.isError) {
    return <div>Error</div>;
  }

  if (query.data === undefined) {
    return <div>Undefined</div>;
  }

  return (
    <ProtectedRoute>
      <ApiTable routes={query.data} />
    </ProtectedRoute>
  );
}

type RouteWithKey = Route & { key: number };

function ApiTable({ routes }: { routes: RouteWithKey[] }) {
  return (
    <TableContainer>
      <Table variant="striped" colorScheme="teal">
        <TableCaption>Api Routes List</TableCaption>
        <Thead>
          <Tr>
            <Th>Route</Th>
            <Th>Method</Th>
            <Th>Path</Th>
            <Th>Params</Th>
          </Tr>
          {routes
            .filter((route) => route.method !== 'HEAD')
            .sort((a, b) => a.path.localeCompare(b.path))
            .map((route) => (
              <ApiRow key={route.key} route={route} />
            ))}
        </Thead>
      </Table>
    </TableContainer>
  );
}

function ApiRow({ route }: { route: Route }) {
  return (
    <Tr>
      <Td>{route.name}</Td>
      <Td>{route.method}</Td>
      <Td>{route.path}</Td>
      <Td>{route.params}</Td>
    </Tr>
  );
}
