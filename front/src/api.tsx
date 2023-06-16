export const globals = {
  protocol: 'http',
  host: 'localhost',
  port: '8080',
  base: 'api',
};

export const endpoints = {
  login: 'auth/login',
  register: 'auth/register',
  user: 'user',
};

export async function getSchema(endpoint: string, schemaName: string) {
  const response = await fetch(
    `${globals.protocol}://${globals.host}:${globals.port}/${globals.base}/${endpoint}/schema/${schemaName}`,
  );
  const jsonData = await response.json();
  return jsonData;
}

export async function sendForm<T>(data: { formData: T }) {
  const response = await fetch('http://localhost:8080/api/auth/register', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data.formData),
  });
  const jsonData: T = await response.json();
  return jsonData;
}
