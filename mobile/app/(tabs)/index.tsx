import React from 'react';
import { View, Text, StyleSheet, ScrollView } from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import { COLORS } from '@/utils/constants';

export default function AboutScreen() {
  return (
    <ScrollView style={styles.container} contentContainerStyle={styles.content}>
      <View style={styles.logoContainer}>
        <Ionicons name="home-outline" size={80} color={COLORS.PRIMARY} />
      </View>
      
      <Text style={styles.title}>About Property Finder</Text>
      
      <Text style={styles.description}>
        Welcome to Property Finder, your trusted companion in discovering the perfect property. 
        Our platform connects you with a curated selection of premium properties, 
        from luxurious villas to cozy apartments.
      </Text>
      
      <Text style={styles.description}>
        Browse through our extensive listings, use our powerful search features, 
        and find detailed information about each property including amenities, 
        pricing, and terms & conditions.
      </Text>
      
      <View style={styles.featureContainer}>
        <View style={styles.feature}>
          <Ionicons name="search" size={24} color={COLORS.PRIMARY} />
          <Text style={styles.featureText}>Smart Search</Text>
        </View>
        
        <View style={styles.feature}>
          <Ionicons name="images" size={24} color={COLORS.PRIMARY} />
          <Text style={styles.featureText}>Photo Gallery</Text>
        </View>
        
        <View style={styles.feature}>
          <Ionicons name="list" size={24} color={COLORS.PRIMARY} />
          <Text style={styles.featureText}>Detailed Info</Text>
        </View>
      </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: COLORS.WHITE,
  },
  content: {
    padding: 24,
    alignItems: 'center',
  },
  logoContainer: {
    marginVertical: 32,
    padding: 24,
    backgroundColor: COLORS.PRIMARY_LIGHT,
    borderRadius: 100,
  },
  title: {
    fontSize: 28,
    fontWeight: '700',
    color: COLORS.GRAY_900,
    marginBottom: 24,
    textAlign: 'center',
  },
  description: {
    fontSize: 16,
    lineHeight: 24,
    color: COLORS.GRAY_600,
    marginBottom: 16,
    textAlign: 'center',
  },
  featureContainer: {
    flexDirection: 'row',
    justifyContent: 'space-around',
    width: '100%',
    marginTop: 32,
  },
  feature: {
    alignItems: 'center',
    flex: 1,
  },
  featureText: {
    marginTop: 8,
    fontSize: 14,
    color: COLORS.GRAY_700,
    textAlign: 'center',
  },
});
