# NGR-Stack (Frontend

)

Este proyecto es una aplicación cliente desarrollada con [React](https://reactjs.org/)

Alunas de las características de este proyecto son:

- Empleo de [TypeScript](https://www.typescriptlang.org/) para el desarrollo de la aplicación.
- Empleo de [React Router](https://reactrouter.com/) para el manejo de rutas.
- Empleo de [React Query](https://react-query.tanstack.com/) para el manejo de datos.
- Generación automática de Formularios a partir de un [JSON Schema](https://json-schema.org/).
- Librería de componentes [Chakra UI](https://chakra-ui.com/).
- Hooks personalizados para el manejo de la autenticación.
- Linting y Formateo automático con [ESLint](https://eslint.org/) y [Prettier](https://prettier.io/).
- Servidor de desarrollo con [Vite](https://vitejs.dev/).

## Tecnologías y Dependencias

- [React](https://reactjs.org/)
- [TypeScript](https://www.typescriptlang.org/)
- [React Router](https://reactrouter.com/)
- [React Query](https://react-query.tanstack.com/)
- [Chakra UI](https://chakra-ui.com/)
- [Vite](https://vitejs.dev/)
- [ESLint](https://eslint.org/)
- [Prettier](https://prettier.io/)
- [JSON Schema](https://json-schema.org/)
- [React JSON Schema Form](https://rjsf-team.github.io/react-jsonschema-form/docs/)
- [Storybook](https://storybook.js.org/)

## Instalación

1. Clonar el repositorio

```bash
git clone XXX
```

2. Instalar las dependencias

Con [npm](https://www.npmjs.com/):

```bash
npm install
```

Con [yarn](https://yarnpkg.com/):

```bash
yarn install
```

Con [pnpm](https://pnpm.js.org/):

```bash
pnpm install
```

3. Iniciar el servidor de desarrollo

Con [npm](https://www.npmjs.com/):

```bash
npm run dev
```

Con [yarn](https://yarnpkg.com/):

```bash
yarn dev
```

Con [pnpm](https://pnpm.js.org/):

```bash
pnpm dev
```

## Paquete `api`

Este proyecto incluye un paquete de utilidades para el manejo de la API. Este paquete se encuentra en la carpeta `src/api` y contiene:

1. Diccionario de rutas de la API:

```ts
export const endpoints: EndpointList = {
  login: 'auth/login',
  register: 'auth/register',
  logout: 'auth/logout',
  currentUser: 'auth/getCurrentUser',
  user: 'user',
  apiRoutes: 'administration/routes',
  ...
};
```

2. Detalles de conexión a la API:

```ts
export const connectionData: ConnectionData = {
  protocol: 'https',
  host: 'localhost',
  port: '8443',
  base: 'api',
};

export function baseApiUrl() {
  return `${connectionData.protocol}://${connectionData.host}:${connectionData.port}/${connectionData.base}`;
}
```

TODO: Cargar los datos de conexión desde un archivo de configuración.

3. Métodos HTTP para el manejo de la API, que además pueden deserializar los datos recibidos:

```ts
async function get<T>(endpoint: string, headers?: Headers): Promise<T> {
  const response = await fetch(`${baseApiUrl()}/${endpoint}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      ...headers,
    },
  });
  const jsonData: T = await response.json();
  return jsonData;
}
```

4. Ficheros con funciones para solicitar datos a la API:

```ts
export async function getUserById(id: string) {
  const headers = {};
  const response = await http.get<ApiResponse<UserDTO>>(endpoints.user + '/' + id, headers);
  return response.data;
}
```

## Paquete `assets`

Contiene los recursos estáticos de la aplicación, como imágenes, iconos, etc.

## Paquete `components`

Contiene los componentes de la aplicación. Estos componentes se dividen en:

- `common`: Componentes comunes a toda la aplicación.
- `user`: Componentes específicos del interfaz para usuarios de la aplicación.
- `admin`: Componentes específicos del interfaz para administradores de la aplicación.

## Paquete `hooks`

Contiene los hooks personalizados de la aplicación. Algunos hooks ya incluidos:

- `useAuth`: Hook para el manejo de la autenticación. Permite iniciar sesión, cerrar sesión, obtener el usuario actual, etc.
- `useLocalStorage`: Hook para el manejo del almacenamiento local del navegador.
- `useTitle`: Hook para el manejo del título de la página.

## Paquete `layouts`

Contiene los layouts de la aplicación. Los layouts son componentes que definen la estructura de la página, como la barra de navegación, el pie de página, etc.

## Paquete `pages`

Contiene las páginas de la aplicación. Las páginas son componentes que definen las rutas de la aplicación.

## Paquete `providers`

Contiene los providers de la aplicación. Los providers son componentes que definen el contexto de la aplicación, como el contexto de autenticación.

## Paquete `rjsf-config`

Contiene la configuración de [React JSON Schema Form](https://rjsf-team.github.io/react-jsonschema-form/docs/). Esta configuración permite generar formularios a partir de un [JSON Schema](https://json-schema.org/).

Actualmente solo contiene el archivo que registra los widgets personalizados y es utilizado en el componente `NgrForm`.

```ts
export const customWidgets: RegistryWidgetsType = {
  PasswordWidget: PasswordInput,
};
```

## Paquete `stories`

Contiene los stories de [Storybook](https://storybook.js.org/). Los stories son componentes que sirven para documentar y probar los componentes de la aplicación.

TODO: Completar la configuración y crear historias para lso componentes.

## Paquete `tests`

Contiene los tests de la aplicación. Los tests sirven para probar los componentes de la aplicación.
