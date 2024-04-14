import { BarcodeScanningResult } from 'expo-camera/next';
import { useState } from 'react';
import { SafeAreaView, StatusBar as RNStatusBar, StyleSheet, Text } from 'react-native';

import { CamScanner } from './CamScanner';

export function ProductScan() {
  const [scannedCode, setScannedCode] = useState(null as string | null);
  if (!scannedCode)
    return (
      <SafeAreaView style={styles.container}>
        <CamScanner
          onBarcodeScanned={(data: BarcodeScanningResult) => {
            if (scannedCode) {
              console.log('Barcode already scanned');
              return;
            }
            setScannedCode(data.data);
          }}
        />
      </SafeAreaView>
    );

  return (
    <SafeAreaView style={styles.container}>
      <Text>Barcode: {scannedCode}</Text>
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
