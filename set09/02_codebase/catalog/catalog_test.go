package catalog

import "testing"

func TestExercises(t *testing.T) {
	if !t.Run("01_FindReturnsStoredBook", testFindReturnsStoredBook) {
		return
	}
	t.Run("02_Search", testSearch)
}

func testFindReturnsStoredBook(t *testing.T) {
	c := New([]Book{{ISBN: "978-0-13", Title: "Go", PriceCents: 3000}})
	book, ok := c.Find("978013")
	if !ok {
		t.Fatal("Find existing book\n  got:  found=false\n  want: found=true")
	}
	book.PriceCents = 2500
	again, _ := c.Find("978013")
	if again.PriceCents != 2500 {
		t.Errorf("updated price\n  got:  %d\n  want: 2500 (Find returned a copy)", again.PriceCents)
	}
}
