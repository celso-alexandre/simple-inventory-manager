import { tokenKey } from '../context/auth/aux';
import { storage } from '../storage';

const baseUrl = 'http://192.168.0.108:3333/api';
export type Product = {
  id: number;
  uuid: string;
  name: string | null;
  barcode: string;
  productGroupId: number | null;
  updatedByUserId: number;
  createdAt: string;
  updatedAt: string;
};

export function handleApiError(response: Response, logout: () => void) {
  if (response.ok) return;
  console.error('API error:', response.status, response.statusText, typeof response.status);
  if (response.status === 401) {
    console.error('Unauthorized. Logging out.');
    logout();
    return alert('Realize o login novamente.');
  }
}

type ProductScanArgs = { type: 'uuid' | 'barcode'; code: string };
export const apiCall = {
  productScan: async ({ code, type }: ProductScanArgs) => {
    const Authorization = (await storage.load({ key: tokenKey })) as string | null;
    console.log('Authorization:', Authorization);
    const url = `${baseUrl}/products-scan`;
    const body = JSON.stringify(type === 'uuid' ? { uuid: code } : { barcode: code });
    console.log(`POST ${url}. Scanned code:`, JSON.stringify(body));
    return fetch(url, { method: 'POST', body, headers: { Authorization: Authorization || '' } });
  },
};
