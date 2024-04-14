import { decode as atob } from 'base-64';
import type { JwtPayload } from 'jsonwebtoken';
import { createContext } from 'react';

import { storage } from '../../storage';

export type User = {
  userId: string;
  username: string;
  // email: string;
};

export const tokenKey = 'inventory-jwt';
export async function login(username: string, password: string, setUser: React.Dispatch<React.SetStateAction<User | null>>) {
  const response = await fetch('http://192.168.0.108:3333/api/login', {
    method: 'POST',
    body: JSON.stringify({ username, password }),
  });
  if (!response.ok) throw new Error('Login failed');
  const body = (await response.json()) as { token: string };
  storage.save({ key: tokenKey, data: body.token });

  console.log('body.token', body.token);
  const decodedToken = decodeJwtToken(body.token);
  if (decodedToken) {
    setUser(decodedToken.user);
  }
}

export function logout(setUser: React.Dispatch<React.SetStateAction<User | null>>) {
  localStorage.removeItem(tokenKey);
  setUser(null);
}

type AuthContextType = {
  user: User | null;
  login: (username: string, password: string) => Promise<void>;
  logout: () => void;
};

export const AuthContext = createContext<AuthContextType>({
  user: null,
  login: () => Promise.resolve(),
  logout: () => {},
});

type TokenData = JwtPayload & {
  user: User;
};
export function decodeJwtToken(token: string) {
  try {
    const decodedToken = JSON.parse(atob(token.split('.')[1])) as TokenData;
    if (!decodedToken?.user) throw new Error('Invalid token');
    if (!decodedToken.user.username || !decodedToken.user.username) throw new Error('Invalid token data');
    if (typeof decodedToken.user.userId !== 'number') throw new Error('Invalid user ID');
    if (typeof decodedToken.user.username !== 'string') throw new Error('Invalid username');
    return decodedToken;
  } catch (error) {
    console.error('Error decoding token:', error);
    return null;
  }
}
