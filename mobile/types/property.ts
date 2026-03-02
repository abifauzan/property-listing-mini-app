export interface Banner {
  url: string;
}

export interface Image {
  url: string;
}

export interface PropertyListing {
  documentId: string;
  Banner: Banner;
  Title: string;
  Price: string;
  createdAt: string;
}

export interface Property extends PropertyListing {
  Description: string;
  Images: Image[];
  Facilities: string[];
  Terms: string;
  Conditions: string;
}

export interface PropertyListingData {
  propertyListings: PropertyListing[];
}

export interface PropertyDetailData {
  propertyListings: Property[];
}

export interface PropertyListingResponse {
  data: PropertyListingData;
}

export interface PropertyDetailResponse {
  data: PropertyDetailData;
}

export interface ErrorResponse {
  error: {
    message: string;
    code?: string;
  };
}
