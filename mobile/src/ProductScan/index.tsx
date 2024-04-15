import { useEffect, useState } from 'react';
import { SafeAreaView, StatusBar as RNStatusBar, StyleSheet, Text, TextInput, View } from 'react-native';

import { Camera } from './Camera';
import { Product, apiCall, handleApiError } from '../api';
import { InputWithLabel } from '../components/InputWithLabel';
import { Select } from '../components/Select';
import { useAuth } from '../context/auth';
import { formatDateTime } from '../util/formatDate';

export function ProductScan() {
  const { user, logout } = useAuth();
  const [scannedCode, setScannedCode] = useState(null as { type: 'uuid' | 'barcode'; code: string } | null);
  const [product, setProduct] = useState(null as Product | null);

  useEffect(() => {
    if (!scannedCode || !user) return;
    console.log('POST /api/products-scan. Scanned code:', scannedCode);
    apiCall.productScan({ ...scannedCode }).then((response) => {
      console.log('POST /api/products-scan. Status:', response.status);
      if (!response.ok) {
        console.error('POST /api/products-scan. Response:', response);
        setScannedCode(null);
        setProduct(null);
        handleApiError(response, logout);
        return alert('Não foi possível consultar o produto.');
      }
      response
        .json()
        .then((data: Product) => {
          console.log('POST /api/products-scan. Response:', data);
          setProduct(data);
        })
        .catch((error) => {
          console.error('POST /api/products-scan. Response:', error);
          setScannedCode(null);
          setProduct(null);
          return alert('Não foi possível consultar o produto.');
        });
    });
  }, [logout, scannedCode, scannedCode?.code, scannedCode?.type, user]);

  if (!scannedCode) {
    return (
      <SafeAreaView style={styles.container}>
        <Camera scannedCode={scannedCode} setScannedCode={setScannedCode} />
      </SafeAreaView>
    );
  }

  return (
    <SafeAreaView style={styles.container}>
      {product && (
        <>
          <Text>ID: {product.uuid}</Text>
          <Text>Cód. Barras: {product.barcode}</Text>
          <Text>Cadastrado em: {formatDateTime(product.createdAt)}</Text>
          <Text>
            Atualizado em: {formatDateTime(product.updatedAt)} por {product.updatedByUserId}
          </Text>
          <View style={styles.divider} />
          <Select
            placeholder={{ label: 'Selecione o grupo...', value: null }}
            value={product.productGroupId}
            onValueChange={(value) => setProduct({ ...product, productGroupId: parseInt(value, 10) || null })}
            items={[{ label: 'padrão', value: 1 }]}
          />
          <InputWithLabel label="Nome" defaultValue={product.name || undefined} onChangeText={(text) => setProduct({ ...product, name: text })} />
        </>
      )}
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    paddingTop: (RNStatusBar.currentHeight || 0) + 20,
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'flex-start',
  },
  divider: {
    width: '100%',
    height: 1,
    backgroundColor: '#ccc',
    marginVertical: 16,
  },
  input: {},
  row: {},
});
