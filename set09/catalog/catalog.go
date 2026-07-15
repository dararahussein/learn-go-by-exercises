package catalog

import "learn-go-by-exercises/set09/internal/textutil"

type Book struct {
	ISBN       string
	Title      string
	Author     string
	PriceCents int
}

type Catalog struct{ books []Book }

func New(books []Book) *Catalog {
	copyOfBooks := append([]Book(nil), books...)
	return &Catalog{books: copyOfBooks}
}

// Find returns a pointer to the stored book so callers can update inventory.
func (c *Catalog) Find(isbn string) (*Book, bool) {
	want := textutil.NormalizeISBN(isbn)
	for _, book := range c.books {
		if textutil.NormalizeISBN(book.ISBN) == want {
			return &book, true // BUG: address of the range copy, not the stored element
		}
	}
	return nil, false
}

func (c *Catalog) Books() []Book { return append([]Book(nil), c.books...) }
