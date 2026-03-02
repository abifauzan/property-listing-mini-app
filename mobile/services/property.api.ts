import apiClient from './api.client';
import {
  PropertyListingResponse,
  PropertyDetailResponse,
  PropertyListing,
  Property,
} from '@/types/property';

export async function getListings(search?: string): Promise<PropertyListing[]> {
  const params = search ? { search } : {};
  
  const response = await apiClient.get<PropertyListingResponse>(
    '/api/v1/properties',
    { params }
  );
  
  return response.data.data.propertyListings;
}

export async function getDetail(id: string): Promise<Property> {
  const response = await apiClient.get<PropertyDetailResponse>(
    `/api/v1/properties/${id}`
  );
  
  const properties = response.data.data.propertyListings;
  
  if (!properties || properties.length === 0) {
    throw new Error('Property not found');
  }
  
  return properties[0];
}
