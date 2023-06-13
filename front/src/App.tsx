import {
  ListItem,
  UnorderedList,
} from "@chakra-ui/react";

import {
  useQueryClient,
  QueryClientProvider,
} from "react-query";

function App() {
  const queryClient = useQueryClient();
  return (
    <QueryClientProvider client={queryClient}>
<UserList />
    </QueryClientProvider>
  );
}

function UserList() {
  return (
    <UnorderedList>
      <ListItem> </ListItem>
    </UnorderedList>
  )
}

export default App;
