import React from 'react';
import { View, Text, TouchableOpacity, StyleSheet, Dimensions } from 'react-native';
import { Image } from 'expo-image';
import { PropertyListing } from '@/types/property';
import { formatPrice } from '@/utils/format';
import { COLORS } from '@/utils/constants';

interface PropertyGridItemProps {
  property: PropertyListing;
  onPress: () => void;
}

const { width } = Dimensions.get('window');
const ITEM_WIDTH = (width - 48) / 2;

export function PropertyGridItem({ property, onPress }: PropertyGridItemProps) {
  return (
    <TouchableOpacity onPress={onPress} style={styles.container} activeOpacity={0.7}>
      <Image
        source={{ uri: property.Banner.url }}
        style={styles.image}
        contentFit="cover"
        transition={200}
      />
      <View style={styles.content}>
        <Text style={styles.title} numberOfLines={2}>
          {property.Title}
        </Text>
        <Text style={styles.price}>{formatPrice(property.Price)}</Text>
      </View>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  container: {
    width: ITEM_WIDTH,
    backgroundColor: COLORS.WHITE,
    borderRadius: 12,
    marginBottom: 16,
    overflow: 'hidden',
    shadowColor: COLORS.BLACK,
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 3,
  },
  image: {
    width: '100%',
    height: 140,
  },
  content: {
    padding: 12,
  },
  title: {
    fontSize: 14,
    fontWeight: '600',
    color: COLORS.GRAY_900,
    marginBottom: 8,
  },
  price: {
    fontSize: 16,
    fontWeight: '700',
    color: COLORS.PRIMARY,
  },
});
