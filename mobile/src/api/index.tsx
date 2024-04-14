const baseUrl = 'http://192.168.0.108:3333/api';
type AuthHeaders = { Authorization: string };

type ProductScanArgs = AuthHeaders & { type: 'uuid' | 'number'; code: string };
export const apiCall = {
  productScan: async ({ code, type, Authorization }: ProductScanArgs) => {
    const url = `${baseUrl}/products-scan`;
    const body = JSON.stringify(type === 'uuid' ? { uuid: code } : { barcode: code });
    console.log(`POST ${url}. Scanned code:`, JSON.stringify(body));
    return fetch(url, { method: 'POST', body, headers: { Authorization } });
  },
};
