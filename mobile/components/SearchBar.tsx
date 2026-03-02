import React, { useState, useCallback } from 'react';
import { View, TextInput, TouchableOpacity, StyleSheet } from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import { debounce } from '@/utils/debounce';
import { DEBOUNCE_DELAY, COLORS } from '@/utils/constants';

interface SearchBarProps {
  onSearch: (query: string) => void;
  placeholder?: string;
}

export function SearchBar({ onSearch, placeholder = 'Search properties...' }: SearchBarProps) {
  const [value, setValue] = useState('');

  const debouncedSearch = useCallback(
    debounce((query: string) => {
      onSearch(query);
    }, DEBOUNCE_DELAY),
    [onSearch]
  );

  const handleChange = (text: string) => {
    setValue(text);
    debouncedSearch(text);
  };

  const handleClear = () => {
    setValue('');
    onSearch('');
  };

  return (
    <View style={styles.container}>
      <Ionicons name="search" size={20} color={COLORS.GRAY_400} style={styles.icon} />
      <TextInput
        style={styles.input}
        value={value}
        onChangeText={handleChange}
        placeholder={placeholder}
        placeholderTextColor={COLORS.GRAY_400}
        autoCapitalize="none"
        autoCorrect={false}
      />
      {value.length > 0 && (
        <TouchableOpacity onPress={handleClear} style={styles.clearButton}>
          <Ionicons name="close-circle" size={20} color={COLORS.GRAY_400} />
        </TouchableOpacity>
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    alignItems: 'center',
    backgroundColor: COLORS.GRAY_100,
    borderRadius: 12,
    paddingHorizontal: 16,
    paddingVertical: 12,
    marginHorizontal: 16,
  },
  icon: {
    marginRight: 8,
  },
  input: {
    flex: 1,
    fontSize: 16,
    color: COLORS.GRAY_900,
  },
  clearButton: {
    padding: 4,
  },
});
