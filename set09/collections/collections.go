package collections

func Filter[T any](items []T, keep func(T) bool) []T {
	var out []T
	for _, item := range items {
		if keep(item) {
			out = append(out, item)
		}
	}
	return out
}

func Map[T, U any](items []T, transform func(T) U) []U {
	return nil // TODO: follow Filter's shape
}
