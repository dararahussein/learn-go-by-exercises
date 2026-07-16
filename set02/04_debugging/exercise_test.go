package debugging

import "testing"

func TestScoreAddChangesOriginal(t *testing.T) {
	s := Score{}
	s.Add(7)
	if got := s.Points(); got != 7 {
		t.Fatalf("after Add(7), Points(): got %d, want 7", got)
	}
}
