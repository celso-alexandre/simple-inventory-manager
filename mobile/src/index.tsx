import { Login } from './Login';
import { ProductScan } from './ProductScan';
import { useAuth } from './context/auth';

export function Main() {
  const { user } = useAuth();

  if (!user) return <Login />;
  return <ProductScan />;
}
