export interface UserPayload {
  name: string;
  password?: string;
}

export interface CreateUserResponse {
  id: number;
}