package playlist

import (
	"slices"
	"testing"
)

func TestExercises(t *testing.T) {
	testPlaylistAdd(t)
	testPlaylistRemove(t)
	testPlaylistTracksReturnsCopy(t)
}

func testPlaylistAdd(t *testing.T) {
	var p Playlist
	p.Add("one")
	p.Add("")
	p.Add("two")
	if !slices.Equal(p.tracks, []string{"one", "two"}) {
		t.Fatalf("stored tracks: got %v, want [one two]", p.tracks)
	}
}

func testPlaylistRemove(t *testing.T) {
	p := Playlist{tracks: []string{"one", "two"}}
	if got := p.Remove(0); !got {
		t.Fatal("Remove(0): got false, want true")
	}
	if got := p.Remove(9); got {
		t.Fatal("Remove(9): got true, want false")
	}
	if !slices.Equal(p.tracks, []string{"two"}) {
		t.Fatalf("stored tracks: got %v, want [two]", p.tracks)
	}
}

func testPlaylistTracksReturnsCopy(t *testing.T) {
	p := Playlist{tracks: []string{"two"}}
	got := p.Tracks()
	if !slices.Equal(got, []string{"two"}) {
		t.Fatalf("Tracks: got %v, want [two]", got)
	}
	got[0] = "changed"
	again := p.Tracks()
	if len(again) != 1 || again[0] != "two" {
		t.Fatalf("Tracks after caller mutation: got %v, want [two]", again)
	}
}
