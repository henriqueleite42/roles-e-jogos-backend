package utils

import "slices"

func Chunkfy[T any](values []T, length int) [][]T {
	var divided [][]T

	chunkSize := (len(values) + length - 1) / length

	for i := 0; i < len(values); i += chunkSize {
		end := i + chunkSize

		if end > len(values) {
			end = len(values)
		}

		divided = append(divided, values[i:end])
	}

	return divided
}

func Diff[T comparable](a, b []T) []T {
	var diff []T
	for _, e := range a {
		if !slices.Contains(b, e) {
			diff = append(diff, e)
		}
	}
	return diff
}
