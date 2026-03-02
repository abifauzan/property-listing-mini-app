import React, { useState } from 'react';
import { View, FlatList, StyleSheet, Dimensions } from 'react-native';
import { Image } from 'expo-image';
import { Image as ImageType } from '@/types/property';
import { COLORS } from '@/utils/constants';

interface ImageCarouselProps {
  images: ImageType[];
}

const { width } = Dimensions.get('window');

export function ImageCarousel({ images }: ImageCarouselProps) {
  const [activeIndex, setActiveIndex] = useState(0);

  const onViewableItemsChanged = React.useRef(({ viewableItems }: any) => {
    if (viewableItems.length > 0) {
      setActiveIndex(viewableItems[0].index || 0);
    }
  }).current;

  const viewabilityConfig = React.useRef({
    itemVisiblePercentThreshold: 50,
  }).current;

  return (
    <View style={styles.container}>
      <FlatList
        data={images}
        horizontal
        pagingEnabled
        showsHorizontalScrollIndicator={false}
        onViewableItemsChanged={onViewableItemsChanged}
        viewabilityConfig={viewabilityConfig}
        keyExtractor={(item, index) => `${item.url}-${index}`}
        renderItem={({ item }) => (
          <Image
            source={{ uri: item.url }}
            style={styles.image}
            contentFit="cover"
            transition={200}
          />
        )}
      />
      <View style={styles.pagination}>
        {images.map((_, index) => (
          <View
            key={index}
            style={[
              styles.dot,
              index === activeIndex && styles.activeDot,
            ]}
          />
        ))}
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    position: 'relative',
  },
  image: {
    width,
    height: 250,
  },
  pagination: {
    position: 'absolute',
    bottom: 16,
    left: 0,
    right: 0,
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
  },
  dot: {
    width: 8,
    height: 8,
    borderRadius: 4,
    backgroundColor: COLORS.WHITE,
    opacity: 0.5,
    marginHorizontal: 4,
  },
  activeDot: {
    opacity: 1,
    width: 24,
  },
});
