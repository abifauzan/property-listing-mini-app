import React, { useState } from 'react';
import { View, FlatList, StyleSheet } from 'react-native';
import { useRouter } from 'expo-router';
import { SearchBar } from '@/components/SearchBar';
import { LayoutToggle } from '@/components/LayoutToggle';
import { PropertyCard } from '@/components/PropertyCard';
import { PropertyGridItem } from '@/components/PropertyGridItem';
import { EmptyState } from '@/components/EmptyState';
import { LoadingSkeleton } from '@/components/LoadingSkeleton';
import { useProperties } from '@/hooks/useProperties';
import { COLORS } from '@/utils/constants';

export default function ListingScreen() {
  const router = useRouter();
  const [search, setSearch] = useState('');
  const [isGrid, setIsGrid] = useState(false);
  
  const { data: properties, isLoading, isError, refetch } = useProperties(search);

  const handlePropertyPress = (id: string) => {
    router.push(`/property/${id}`);
  };

  return (
    <View style={styles.container}>
      <View style={styles.header}>
        <View style={styles.searchContainer}>
          <SearchBar onSearch={setSearch} />
        </View>
        <LayoutToggle isGrid={isGrid} onToggle={() => setIsGrid(!isGrid)} />
      </View>

      {isLoading ? (
        <LoadingSkeleton isGrid={isGrid} />
      ) : isError ? (
        <EmptyState message="Failed to load properties" icon="alert-circle-outline" />
      ) : !properties || properties.length === 0 ? (
        <EmptyState message="No properties found" />
      ) : (
        <FlatList
          data={properties}
          key={isGrid ? 'grid' : 'list'}
          numColumns={isGrid ? 2 : 1}
          columnWrapperStyle={isGrid ? styles.gridRow : undefined}
          contentContainerStyle={styles.listContent}
          keyExtractor={(item) => item.documentId}
          renderItem={({ item }) =>
            isGrid ? (
              <PropertyGridItem
                property={item}
                onPress={() => handlePropertyPress(item.documentId)}
              />
            ) : (
              <PropertyCard
                property={item}
                onPress={() => handlePropertyPress(item.documentId)}
              />
            )
          }
          refreshing={isLoading}
          onRefresh={refetch}
        />
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: COLORS.WHITE,
  },
  header: {
    flexDirection: 'row',
    alignItems: 'center',
    paddingVertical: 12,
    backgroundColor: COLORS.WHITE,
    borderBottomWidth: 1,
    borderBottomColor: COLORS.GRAY_200,
  },
  searchContainer: {
    flex: 1,
  },
  listContent: {
    paddingVertical: 8,
  },
  gridRow: {
    justifyContent: 'space-between',
    paddingHorizontal: 16,
  },
});
