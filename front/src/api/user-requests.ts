import { ApiResponse } from '../types/api-response';
import { UserDTO } from '../types/generated/models';
import { endpoints } from './endpoints';
import { http } from './http-methods';

export async function getUserById(id: string) {
  const headers = {};
  const response = await http.get<ApiResponse<UserDTO>>(endpoints.user + '/' + id, headers);
  return response.data;
}
