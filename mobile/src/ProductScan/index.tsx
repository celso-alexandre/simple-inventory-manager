import { useEffect, useState } from 'react';
import { SafeAreaView, StatusBar as RNStatusBar, StyleSheet, Text } from 'react-native';

import { Camera } from './Camera';
import { apiCall } from '../api';
import { useAuth } from '../context/auth';

export function ProductScan() {
  const [scannedCode, setScannedCode] = useState(null as { type: 'uuid' | 'number'; code: string } | null);
  const { user } = useAuth();

  useEffect(() => {
    if (!scannedCode) return;
    console.log('POST /api/products-scan. Scanned code:', scannedCode);
    apiCall.productScan({ ...scannedCode, Authorization: user!.Authorization }).then((response) => {
      console.log('POST /api/products-scan. Status:', response.status);
      if (!response.ok) return console.error('POST /api/products-scan. Error:', response.statusText);
      response.json().then((data) => {
        console.log('POST /api/products-scan. Response:', data);
      });
    });
  }, [scannedCode, scannedCode?.code, scannedCode?.type, user]);

  return (
    <SafeAreaView style={styles.container}>
      {!scannedCode ? (
        <Camera scannedCode={scannedCode} setScannedCode={setScannedCode} />
      ) : (
        <Text>
          {scannedCode.type === 'uuid' ? 'Código' : 'Cód. Barras'}: {scannedCode.code}
        </Text>
      )}
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    paddingTop: RNStatusBar.currentHeight,
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'flex-start',
  },
  text: {},
});
