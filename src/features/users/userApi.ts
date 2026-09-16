import { api } from '../../components/api';
import { UserPayload, CreateUserResponse } from './types';

export const userApi = {
  create: async (data: UserPayload) => {
    const res = await api.post<CreateUserResponse>('/users', data);
    return res.data;
  },
  delete: async (data: UserPayload) => {
    const res = await api.delete<{ message: string }>('/users', { data });
    return res.data;
  },
};