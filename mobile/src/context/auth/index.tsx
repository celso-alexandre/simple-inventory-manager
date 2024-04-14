import React, { useState, useEffect } from 'react';

import { AuthContext, User, decodeJwtToken, login, logout, tokenKey } from './aux';
import { storage } from '../../storage';

type IProps = {
  children: React.ReactNode;
};
export function AuthProvider({ children }: IProps) {
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    storage.load({ key: tokenKey }).then((token) => {
      if (token) {
        const decodedToken = decodeJwtToken(token);
        setUser(decodedToken?.user || null);
      }
    });
  }, []);

  async function loginFn(username: string, password: string) {
    await login(username, password, setUser);
  }
  function logoutFn() {
    logout(setUser);
  }
  return <AuthContext.Provider value={{ user, login: loginFn, logout: logoutFn }}>{children}</AuthContext.Provider>;
}

export function useAuth() {
  return React.useContext(AuthContext);
}
