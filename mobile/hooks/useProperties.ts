import { useQuery } from '@tanstack/react-query';
import { getListings } from '@/services/property.api';
import { QUERY_KEYS } from '@/utils/constants';

export function useProperties(search?: string) {
  return useQuery({
    queryKey: [QUERY_KEYS.PROPERTIES, search],
    queryFn: () => getListings(search),
    staleTime: 5 * 60 * 1000,
  });
}
