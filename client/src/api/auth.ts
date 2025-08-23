import { API_URL } from '@/config';
import axios from 'axios';

export interface User {
  email: string;
  full_name: string;
  avatar_url: string;
}

export interface SignInResponse {
  access_token: string;
  user: User;
}

const apiClient = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const signInWithGoogle = async (idToken: string): Promise<SignInResponse> => {
  try {
    const response = await apiClient.post<SignInResponse>('/auth/signin', {
      id_token: idToken,
    });
    return response.data;
  } catch (error) {
    if (axios.isAxiosError(error) && error.response) {
      throw new Error(error.response.data.error || 'An unknown error occurred');
    }
    throw error;
  }
};
