package slice

func Remove[T any](items []T, index int) []T {
	items[index] = items[len(items)-1]

	return items[:len(items)-1]
}
