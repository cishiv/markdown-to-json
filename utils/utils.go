package utils

/**
Map: Define a generic map function that accepts a slice<T> and a function(T): V,
that applies f(T):V to each element of slice<T> such that slice<T> is transformed to slice<V>
**/
func Map[T any, R any](operand []T, f func(T) R) []R {
	result := make([]R, len(operand))
	for i, r := range operand {
		result[i] = f(r)
	}
	return result
}

/**
Contains: Assert whether slice<T> contains e<T>
**/
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
