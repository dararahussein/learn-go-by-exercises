package playlist

import (
	"slices"
	"testing"
)

func TestExercises(t *testing.T) {
	if !t.Run("01_Add", testPlaylistAdd) {
		return
	}
	if !t.Run("02_Remove", testPlaylistRemove) {
		return
	}
	t.Run("03_TracksReturnsCopy", testPlaylistTracksReturnsCopy)
}

func testPlaylistAdd(t *testing.T) {
	var p Playlist
	p.Add("one")
	p.Add("")
	p.Add("two")
	if !slices.Equal(p.tracks, []string{"one", "two"}) {
		t.Errorf("stored tracks = %v; want [one two]", p.tracks)
	}
}

func testPlaylistRemove(t *testing.T) {
	p := Playlist{tracks: []string{"one", "two"}}
	if !p.Remove(0) || p.Remove(9) {
		t.Fatal("Remove should accept 0 and reject 9")
	}
	if !slices.Equal(p.tracks, []string{"two"}) {
		t.Errorf("stored tracks = %v; want [two]", p.tracks)
	}
}

func testPlaylistTracksReturnsCopy(t *testing.T) {
	p := Playlist{tracks: []string{"two"}}
	got := p.Tracks()
	if !slices.Equal(got, []string{"two"}) {
		t.Fatalf("Tracks = %v; want [two]", got)
	}
	got[0] = "changed"
	again := p.Tracks()
	if len(again) != 1 || again[0] != "two" {
		t.Error("Tracks exposed Playlist's internal slice")
	}
}
