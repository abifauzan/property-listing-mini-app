import React from 'react';
import { View, Text, StyleSheet, ScrollView, TouchableOpacity, Alert } from 'react-native';
import { useLocalSearchParams } from 'expo-router';
import { Image } from 'expo-image';
import { ImageCarousel } from '@/components/ImageCarousel';
import { LoadingSkeleton } from '@/components/LoadingSkeleton';
import { EmptyState } from '@/components/EmptyState';
import { usePropertyDetail } from '@/hooks/usePropertyDetail';
import { formatPrice } from '@/utils/format';
import { COLORS } from '@/utils/constants';

export default function PropertyDetailScreen() {
  const { id } = useLocalSearchParams<{ id: string }>();
  const { data: property, isLoading, isError } = usePropertyDetail(id);

  const handleBookNow = () => {
    Alert.alert(
      'Book Property',
      `You are about to book "${property?.Title}". This is a demo action.`,
      [
        { text: 'Cancel', style: 'cancel' },
        { text: 'Confirm', onPress: () => Alert.alert('Success', 'Booking confirmed!') },
      ]
    );
  };

  if (isLoading) {
    return (
      <View style={styles.container}>
        <LoadingSkeleton />
      </View>
    );
  }

  if (isError || !property) {
    return (
      <View style={styles.container}>
        <EmptyState message="Property not found" icon="alert-circle-outline" />
      </View>
    );
  }

  return (
    <View style={styles.container}>
      <ScrollView style={styles.scrollView}>
        <Image
          source={{ uri: property.Banner.url }}
          style={styles.heroImage}
          contentFit="cover"
          transition={200}
        />

        <View style={styles.content}>
          <Text style={styles.title}>{property.Title}</Text>
          <Text style={styles.price}>{formatPrice(property.Price)}</Text>

          {property.Images && property.Images.length > 0 && (
            <View style={styles.section}>
              <Text style={styles.sectionTitle}>Gallery</Text>
              <ImageCarousel images={property.Images} />
            </View>
          )}

          <View style={styles.section}>
            <Text style={styles.sectionTitle}>Description</Text>
            <Text style={styles.description}>{property.Description}</Text>
          </View>

          {property.Facilities && property.Facilities.length > 0 && (
            <View style={styles.section}>
              <Text style={styles.sectionTitle}>Facilities</Text>
              <View style={styles.facilitiesContainer}>
                {property.Facilities.map((facility, index) => (
                  <View key={index} style={styles.facilityChip}>
                    <Text style={styles.facilityText}>{facility}</Text>
                  </View>
                ))}
              </View>
            </View>
          )}

          {property.Terms && (
            <View style={styles.section}>
              <Text style={styles.sectionTitle}>Terms & Conditions</Text>
              <Text style={styles.description}>{property.Terms}</Text>
            </View>
          )}

          {property.Conditions && (
            <View style={styles.section}>
              <Text style={styles.sectionTitle}>Additional Conditions</Text>
              <Text style={styles.description}>{property.Conditions}</Text>
            </View>
          )}

          <View style={styles.bottomSpacer} />
        </View>
      </ScrollView>

      <View style={styles.footer}>
        <TouchableOpacity style={styles.bookButton} onPress={handleBookNow}>
          <Text style={styles.bookButtonText}>Book Now</Text>
        </TouchableOpacity>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: COLORS.WHITE,
  },
  scrollView: {
    flex: 1,
  },
  heroImage: {
    width: '100%',
    height: 300,
  },
  content: {
    padding: 16,
  },
  title: {
    fontSize: 24,
    fontWeight: '700',
    color: COLORS.GRAY_900,
    marginBottom: 8,
  },
  price: {
    fontSize: 28,
    fontWeight: '700',
    color: COLORS.PRIMARY,
    marginBottom: 24,
  },
  section: {
    marginBottom: 24,
  },
  sectionTitle: {
    fontSize: 18,
    fontWeight: '600',
    color: COLORS.GRAY_900,
    marginBottom: 12,
  },
  description: {
    fontSize: 15,
    lineHeight: 22,
    color: COLORS.GRAY_600,
  },
  facilitiesContainer: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    gap: 8,
  },
  facilityChip: {
    backgroundColor: COLORS.PRIMARY_LIGHT,
    paddingHorizontal: 16,
    paddingVertical: 8,
    borderRadius: 20,
  },
  facilityText: {
    fontSize: 14,
    color: COLORS.PRIMARY,
    fontWeight: '500',
  },
  bottomSpacer: {
    height: 80,
  },
  footer: {
    position: 'absolute',
    bottom: 0,
    left: 0,
    right: 0,
    padding: 16,
    backgroundColor: COLORS.WHITE,
    borderTopWidth: 1,
    borderTopColor: COLORS.GRAY_200,
    shadowColor: COLORS.BLACK,
    shadowOffset: { width: 0, height: -2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 5,
  },
  bookButton: {
    backgroundColor: COLORS.PRIMARY,
    paddingVertical: 16,
    borderRadius: 12,
    alignItems: 'center',
  },
  bookButtonText: {
    color: COLORS.WHITE,
    fontSize: 18,
    fontWeight: '600',
  },
});
