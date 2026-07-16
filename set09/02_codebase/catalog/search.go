package catalog

// Search returns books whose title or author contains query, ignoring case.
// Results preserve catalog order. An empty query returns no books.
func (c *Catalog) Search(query string) []Book {
	return nil // TODO: write tests first
}
