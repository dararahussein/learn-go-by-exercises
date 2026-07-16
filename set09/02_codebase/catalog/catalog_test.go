package catalog

import "testing"

func TestExercises(t *testing.T) {
	testFindReturnsStoredBook(t)
	testSearch(t)
}

func testFindReturnsStoredBook(t *testing.T) {
	c := New([]Book{{ISBN: "978-0-13", Title: "Go", PriceCents: 3000}})
	book, ok := c.Find("978013")
	if !ok {
		t.Fatal("Find existing book: got found=false, want found=true")
	}
	book.PriceCents = 2500
	again, _ := c.Find("978013")
	if again.PriceCents != 2500 {
		t.Fatalf("updated price: got %d, want 2500 (Find returned a copy)", again.PriceCents)
	}
}
