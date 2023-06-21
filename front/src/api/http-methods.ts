import { baseApiUrl } from './connection-data';

type Headers = {
  [key: string]: string;
};

/**
 * Fetches data from the server using the GET method
 *
 * The endpoint is appended to the base url
 * List of endpoints can be found on {@link ./endpoints.ts}
 * Uses the {@link ./connection-data.ts} to get the base url
 *
 * @param endpoint
 * @param {Headers} headers
 * @returns {T} The response from the server
 */
async function get<T>(endpoint: string, headers?: Headers) {
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

/**
 * Sends data to the server using the POST method
 *
 * The endpoint is appended to the base url
 * List of endpoints can be found on {@link ./endpoints.ts}
 * Uses the {@link ./connection-data.ts} to get the base url
 *
 * @param endpoint
 * @param {Req} data
 * @param {Headers} headers
 * @returns {Res} The response from the server
 */
async function post<Req, Res>(endpoint: string, data: Req, headers?: Headers) {
  const response = await fetch(`${baseApiUrl()}/${endpoint}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      ...headers,
    },
    body: JSON.stringify(data),
  });
  const jsonData: Res = await response.json();
  return jsonData;
}

/**
 * Sends data to the server using the PUT method
 *
 * The endpoint is appended to the base url
 * List of endpoints can be found on {@link ./endpoints.ts}
 * Uses the {@link ./connection-data.ts} to get the base url
 *
 * @param endpoint
 * @param {Req} data
 * @param {Headers} headers
 * @returns {Res} The response from the server
 */
async function put<Req, Res>(endpoint: string, data: Req, headers?: Headers) {
  const response = await fetch(`${baseApiUrl()}/${endpoint}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      ...headers,
    },
    body: JSON.stringify(data),
  });
  const jsonData: Res = await response.json();
  return jsonData;
}

/**
 * Sends data to the server using the DELETE method
 *
 * The endpoint is appended to the base url
 * List of endpoints can be found on {@link ./endpoints.ts}
 * Uses the {@link ./connection-data.ts} to get the base url
 *
 * @param endpoint
 * @param {Headers} headers
 * @returns {Res} The response from the server
 */
async function del<Res>(endpoint: string, headers?: Headers) {
  const response = await fetch(`${baseApiUrl()}/${endpoint}`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
      ...headers,
    },
  });
  const jsonData: Res = await response.json();
  return jsonData;
}

export const http = {
  get,
  post,
  put,
  del,
};
