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
		t.Errorf("stored tracks\n  got:  %v\n  want: [one two]", p.tracks)
	}
}

func testPlaylistRemove(t *testing.T) {
	p := Playlist{tracks: []string{"one", "two"}}
	if got := p.Remove(0); !got {
		t.Fatal("Remove(0)\n  got:  false\n  want: true")
	}
	if got := p.Remove(9); got {
		t.Fatal("Remove(9)\n  got:  true\n  want: false")
	}
	if !slices.Equal(p.tracks, []string{"two"}) {
		t.Errorf("stored tracks\n  got:  %v\n  want: [two]", p.tracks)
	}
}

func testPlaylistTracksReturnsCopy(t *testing.T) {
	p := Playlist{tracks: []string{"two"}}
	got := p.Tracks()
	if !slices.Equal(got, []string{"two"}) {
		t.Fatalf("Tracks\n  got:  %v\n  want: [two]", got)
	}
	got[0] = "changed"
	again := p.Tracks()
	if len(again) != 1 || again[0] != "two" {
		t.Errorf("Tracks after caller mutation\n  got:  %v\n  want: [two]", again)
	}
}
