export function formatPrice(price: string | number): string {
  const numPrice = typeof price === 'string' ? parseInt(price, 10) : price;
  
  if (isNaN(numPrice)) {
    return 'Rp 0';
  }
  
  return `Rp ${numPrice.toLocaleString('id-ID')}`;
}

export function formatDate(dateString: string): string {
  try {
    const date = new Date(dateString);
    return date.toLocaleDateString('id-ID', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
    });
  } catch {
    return dateString;
  }
}
