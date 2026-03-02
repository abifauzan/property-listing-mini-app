import React from 'react';
import { TouchableOpacity, StyleSheet } from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import { COLORS } from '@/utils/constants';

interface LayoutToggleProps {
  isGrid: boolean;
  onToggle: () => void;
}

export function LayoutToggle({ isGrid, onToggle }: LayoutToggleProps) {
  return (
    <TouchableOpacity onPress={onToggle} style={styles.button}>
      <Ionicons
        name={isGrid ? 'list' : 'grid'}
        size={24}
        color={COLORS.PRIMARY}
      />
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  button: {
    padding: 8,
    marginRight: 16,
  },
});
