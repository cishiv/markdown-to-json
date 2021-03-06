package utils

import (
	"encoding/json"
	"strings"
)

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

/**
ContainsAny: Does slice 2 contain any of the elements in slice 1
s1: Slice<T>
s2: Slice<T>
**/
func ContainsAny[T comparable](s1 []T, s2 []T) bool {
	doesContain := false
	for _, e := range s2 {
		// if Contains, flip the bit
		doesContain = !Contains(s1, e)
		if !doesContain {
			break
		}
	}
	return !doesContain
}

/**
Difference: Compute the set difference between s1 & s2
**/
func Difference[T comparable](s1, s2 []T) (difference []T) {
	unset := make(map[T]bool)
	for _, item := range s2 {
		unset[item] = true
	}
	for _, item := range s1 {
		if _, q := unset[item]; !q {
			difference = append(difference, item)
		}
	}
	return difference
}

/**
MapToJsonString: Return a string representation of a <K,V> map
**/
func MapToJsonString[K comparable, V any](m map[K]V) string {
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(b)
}

/**
TrimAndCount: Trim the left most whitespaces in a string and return
(trimmedString, spacesTrimmed)
**/
func TrimAndCount(s string) (string, int) {
	spaces := len(s) - len(strings.TrimLeft(s, " "))
	trimmed := strings.TrimLeft(s, " ")
	return trimmed, spaces
}

func Matrix2D[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		endRow := startRow + m
		matrix[i] = rows[startRow:endRow:endRow]
	}
	return matrix
}
