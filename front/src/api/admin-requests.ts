import { ApiResponse } from '../types/api-response';
import { Route } from '../types/generated/models';
import { endpoints } from './endpoints';
import { http } from './http-methods';

export async function getApiRoutes(token: string) {
  const headers = {
    Authorization: `Bearer ${token}`,
  };
  const response = await http.get<ApiResponse<Route[]>>(endpoints.apiRoutes, headers);
  return response.data;
}
