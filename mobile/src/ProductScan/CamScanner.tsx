import { requestCameraPermissionsAsync } from 'expo-camera';
import { CameraProps, CameraView } from 'expo-camera/next';
import { useEffect, useState } from 'react';
import { Button, StyleSheet, Text } from 'react-native';

export function CamScanner(props: CameraProps) {
  const [hasPermission, setHasPermission] = useState(null as boolean | null);
  const [scanned, setScanned] = useState(false);

  useEffect(() => {
    (async () => {
      const { status } = await requestCameraPermissionsAsync();
      await new Promise((resolve) => setTimeout(resolve, 1000));
      setHasPermission(status === 'granted');
    })();
  }, []);

  if (hasPermission === null) {
    return <Text>Solicitando acesso à câmera</Text>;
  }
  if (hasPermission === false) {
    return <Button title="Autorize o acesso à câmera" onPress={() => requestCameraPermissionsAsync()} />;
  }

  return (
    <CameraView
      barcodeScannerSettings={{
        barcodeTypes: ['aztec', 'ean13', 'ean8', 'qr', 'pdf417', 'upc_e', 'datamatrix', 'code39', 'code93', 'itf14', 'codabar', 'code128', 'upc_a'],
      }}
      style={[StyleSheet.absoluteFill, styles.container]}
      {...props}
      onBarcodeScanned={
        scanned
          ? undefined
          : async (data) => {
              props.onBarcodeScanned?.(data);
              await new Promise((resolve) => setTimeout(resolve, 1000));
              setScanned(true);
              // console.log('Barcode data:', data);
              // alert(`Bar code with type ${data.type} and data ${data.data} has been scanned!`);
            }
      }
    >
      {/* <Text style={styles.description}>Scan your QR code</Text> */}
      {/* <Text onPress={() => navigation.navigate('TempRecordStack')} style={styles.cancel}>
          Back
        </Text> */}
      {/* {scanned && <Button title="Tap to Scan Again" onPress={() => setScanned(false)} />} */}
    </CameraView>
  );
}

// const { width } = Dimensions.get('window');
const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
  },
  // description: {
  //   fontSize: width * 0.09,
  //   fontWeight: 'bold',
  //   marginTop: '30%',
  //   textAlign: 'center',
  //   width: '70%',
  //   //  color: colors.purple,
  // },
  // cancel: {
  //   //  backgroundColor: colors.dark,
  //   fontSize: width * 0.05,
  //   fontWeight: 'bold',
  //   padding: 4,
  //   textAlign: 'center',
  //   width: '30%',
  //   opacity: 0.6,
  //   //  color: colors.purple,
  // },
});
