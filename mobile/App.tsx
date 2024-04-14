import { Main } from './src';
import { AuthProvider } from './src/context/auth';

export default function App() {
  return (
    <AuthProvider>
      <Main />
    </AuthProvider>
  );
}
