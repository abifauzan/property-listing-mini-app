import React from 'react';
import { View, Text, TouchableOpacity, StyleSheet, Dimensions } from 'react-native';
import { Image } from 'expo-image';
import { PropertyListing } from '@/types/property';
import { formatPrice } from '@/utils/format';
import { COLORS } from '@/utils/constants';

interface PropertyCardProps {
  property: PropertyListing;
  onPress: () => void;
}

const { width } = Dimensions.get('window');

export function PropertyCard({ property, onPress }: PropertyCardProps) {
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
    flexDirection: 'row',
    backgroundColor: COLORS.WHITE,
    borderRadius: 12,
    marginHorizontal: 16,
    marginVertical: 8,
    overflow: 'hidden',
    shadowColor: COLORS.BLACK,
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 3,
  },
  image: {
    width: 120,
    height: 120,
  },
  content: {
    flex: 1,
    padding: 12,
    justifyContent: 'space-between',
  },
  title: {
    fontSize: 16,
    fontWeight: '600',
    color: COLORS.GRAY_900,
    marginBottom: 8,
  },
  price: {
    fontSize: 18,
    fontWeight: '700',
    color: COLORS.PRIMARY,
  },
});
