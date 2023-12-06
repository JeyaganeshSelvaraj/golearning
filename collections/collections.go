package generics

func addToSlice[T any](student T, lst []T) []T {
	return append(lst, student)
}

func GenericFilter[T any](lst []T, fn func(item T) bool) []T {
	var result []T
	for _, it := range lst {
		if fn(it) {
			result = append(result, it)
		}
	}
	return result
}
func GenericMap[T1, T2 any](lst []T1, fn func(item T1) T2) []T2 {
	var result []T2
	for _, it := range lst {
		result = append(result, fn(it))
	}
	return result
}
