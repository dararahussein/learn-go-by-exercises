package set02

import (
	"slices"
	"testing"
)

func TestPlaylist(t *testing.T) {
	var p Playlist
	p.Add("one")
	p.Add("")
	p.Add("two")
	if !p.Remove(0) || p.Remove(9) {
		t.Error("Remove should accept 0 and reject 9")
	}
	got := p.Tracks()
	if !slices.Equal(got, []string{"two"}) {
		t.Errorf("Tracks = %v; want [two]", got)
	}
	got[0] = "changed"
	if p.Tracks()[0] != "two" {
		t.Error("Tracks exposed Playlist's internal slice")
	}
}
