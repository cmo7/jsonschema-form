import { ApiResponse } from '../types/api-response';
import { UserResponse } from '../types/generated/models';
import { endpoints } from './endpoints';
import { http } from './http-methods';



export async function sendForm<Req, Res>(endpoint: string, data: Req) {
  console.log('sendForm', endpoint, data);
  return http.post<Req, Res>(endpoint, data);
}

export async function logout() {
  return http.post<unknown, unknown>(endpoints.logout, {});
}

export async function getCurrentUser(token: string) {
  const headers = {
    Authorization: `Bearer ${token}`,
  };
  const response = await http.get<ApiResponse<UserResponse>>(endpoints.currentUser, headers);
  return response.data;
}
