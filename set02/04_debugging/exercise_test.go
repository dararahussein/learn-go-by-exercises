package debugging

import "testing"

func TestScoreAddChangesOriginal(t *testing.T) {
	s := Score{}
	s.Add(7)
	if got := s.Points(); got != 7 {
		t.Errorf("after Add(7), Points()\n  got:  %d\n  want: 7", got)
	}
}
