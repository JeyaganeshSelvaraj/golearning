package generics

func addToSlice[T any](student T, lst []T) []T {
	return append(lst, student)
}
