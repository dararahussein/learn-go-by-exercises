package pricing

import "learn-go-by-exercises/set09/catalog"

// FinalPrice applies a whole-number discount percentage.
func FinalPrice(book catalog.Book, discountPercent int) int {
	if discountPercent < 0 {
		discountPercent = 0
	}
	if discountPercent > 100 {
		discountPercent = 100
	}
	return book.PriceCents * (100 - discountPercent) / 100
}
