import { useQuery } from '@tanstack/react-query';
import { getDetail } from '@/services/property.api';
import { QUERY_KEYS } from '@/utils/constants';

export function usePropertyDetail(id: string) {
  return useQuery({
    queryKey: [QUERY_KEYS.PROPERTY_DETAIL, id],
    queryFn: () => getDetail(id),
    enabled: !!id,
    staleTime: 10 * 60 * 1000,
  });
}
