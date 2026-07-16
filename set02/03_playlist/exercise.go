package playlist

// Playlist is the end-of-set application: a struct that owns a slice and uses
// pointer-receiver methods to mutate it.
type Playlist struct{ tracks []string }

func (p *Playlist) Add(track string) {
	// TODO: ignore empty tracks; append other tracks
}

func (p *Playlist) Remove(index int) bool {
	return false // TODO: remove a valid index and report success
}

func (p Playlist) Tracks() []string {
	return nil // TODO: return a copy so callers cannot mutate internal state
}
