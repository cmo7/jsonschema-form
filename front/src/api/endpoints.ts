type EndpointList = {
  [key: string]: string;
};

export const endpoints: EndpointList = {
  login: 'auth/login',
  register: 'auth/register',
  logout: 'auth/logout',
  currentUser: 'auth/getCurrentUser',
  user: 'user',
  apiRoutes: 'administration/routes',
};
