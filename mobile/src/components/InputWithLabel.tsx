import React from 'react';
import { View, TextInput, Text, StyleSheet, TextInputProps } from 'react-native';

type IProps = TextInputProps & {
  label: string;
};
export function InputWithLabel({ label, value, onChangeText, ...props }: IProps) {
  return (
    <View style={styles.container}>
      <Text style={styles.label}>{label}</Text>
      <TextInput style={styles.input} value={value} onChangeText={onChangeText} {...props} />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    marginBottom: 16,
  },
  label: {
    fontSize: 16,
    marginRight: 8,
  },
  input: {
    borderWidth: 1,
    borderColor: '#ccc',
    borderRadius: 5,
    paddingTop: 1,
    paddingBottom: 1,
    paddingLeft: 10,
    fontSize: 16,
    width: 200, // Adjust width as needed
  },
});
