import { ChakraProvider } from "@chakra-ui/react";
import React from "react";
import ReactDOM from "react-dom/client";

import {
  RouterProvider,
  createBrowserRouter
} from "react-router-dom";
import Root from "./routes/root";
import UserProfile from "./routes/user-profile";
import UserList from "./routes/user-list";
import CreateUser from "./routes/create-user";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
  },
  {
    path: "/users/:userId",
    element: <UserProfile />,
  },
  {
    path: "/users-list",
    element: <UserList />,
  },
  {
    path: "/create-user",
    element: <CreateUser />,
  }

])

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <ChakraProvider>
      <RouterProvider router={router} />
    </ChakraProvider>
  </React.StrictMode>
);
