import React, { useEffect } from 'react';
import { View, StyleSheet, Dimensions } from 'react-native';
import Animated, {
  useSharedValue,
  useAnimatedStyle,
  withRepeat,
  withTiming,
  interpolate,
} from 'react-native-reanimated';
import { COLORS } from '@/utils/constants';

const { width } = Dimensions.get('window');

interface LoadingSkeletonProps {
  isGrid?: boolean;
}

export function LoadingSkeleton({ isGrid = false }: LoadingSkeletonProps) {
  const opacity = useSharedValue(0.3);

  useEffect(() => {
    opacity.value = withRepeat(
      withTiming(1, { duration: 1000 }),
      -1,
      true
    );
  }, []);

  const animatedStyle = useAnimatedStyle(() => ({
    opacity: opacity.value,
  }));

  if (isGrid) {
    return (
      <View style={styles.gridContainer}>
        {[1, 2, 3, 4].map((item) => (
          <Animated.View key={item} style={[styles.gridItem, animatedStyle]}>
            <View style={styles.gridImage} />
            <View style={styles.gridContent}>
              <View style={styles.gridTitle} />
              <View style={styles.gridPrice} />
            </View>
          </Animated.View>
        ))}
      </View>
    );
  }

  return (
    <View style={styles.listContainer}>
      {[1, 2, 3].map((item) => (
        <Animated.View key={item} style={[styles.listItem, animatedStyle]}>
          <View style={styles.listImage} />
          <View style={styles.listContent}>
            <View style={styles.listTitle} />
            <View style={styles.listPrice} />
          </View>
        </Animated.View>
      ))}
    </View>
  );
}

const ITEM_WIDTH = (width - 48) / 2;

const styles = StyleSheet.create({
  gridContainer: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    justifyContent: 'space-between',
    paddingHorizontal: 16,
  },
  gridItem: {
    width: ITEM_WIDTH,
    backgroundColor: COLORS.GRAY_100,
    borderRadius: 12,
    marginBottom: 16,
    overflow: 'hidden',
  },
  gridImage: {
    width: '100%',
    height: 140,
    backgroundColor: COLORS.GRAY_200,
  },
  gridContent: {
    padding: 12,
  },
  gridTitle: {
    height: 16,
    backgroundColor: COLORS.GRAY_200,
    borderRadius: 4,
    marginBottom: 8,
  },
  gridPrice: {
    height: 20,
    width: '60%',
    backgroundColor: COLORS.GRAY_200,
    borderRadius: 4,
  },
  listContainer: {
    paddingVertical: 8,
  },
  listItem: {
    flexDirection: 'row',
    backgroundColor: COLORS.GRAY_100,
    borderRadius: 12,
    marginHorizontal: 16,
    marginVertical: 8,
    overflow: 'hidden',
  },
  listImage: {
    width: 120,
    height: 120,
    backgroundColor: COLORS.GRAY_200,
  },
  listContent: {
    flex: 1,
    padding: 12,
    justifyContent: 'space-between',
  },
  listTitle: {
    height: 16,
    backgroundColor: COLORS.GRAY_200,
    borderRadius: 4,
    marginBottom: 8,
  },
  listPrice: {
    height: 24,
    width: '50%',
    backgroundColor: COLORS.GRAY_200,
    borderRadius: 4,
  },
});
