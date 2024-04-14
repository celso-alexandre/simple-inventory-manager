import { BarcodeScanningResult } from 'expo-camera/next';

import { CamScanner } from '../components/CamScanner';

type ScannedCode = { type: 'uuid' | 'number'; code: string } | null;
type IProps = {
  scannedCode: ScannedCode;
  setScannedCode: React.Dispatch<React.SetStateAction<ScannedCode>>;
};
export function Camera({ scannedCode, setScannedCode }: IProps) {
  return (
    <CamScanner
      enableTorch
      onBarcodeScanned={(data: BarcodeScanningResult) => {
        if (scannedCode) {
          console.log('Barcode already scanned');
          return;
        }
        const code = data.data;
        console.log('Barcode data:', code);
        if (!code || typeof code !== 'string') {
          console.log('Invalid barcode data (not a string):', code);
          return;
        }
        const isNumber = code.match(/^\d+$/);
        const isUuid = code.match(/^[0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12}$/i);
        if (!isNumber && !isUuid) {
          console.log('Invalid barcode: (not a valid number regex or uuid)', code);
          return;
        }
        setScannedCode({ type: isUuid ? 'uuid' : 'number', code });
      }}
    />
  );
}
