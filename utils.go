package main

func isArraysEquals(arr []string, arr2 []string) bool {
	if len(arr) != len(arr2) {
		return false
	}
	for i, val := range arr {
		if arr2[i] != val {
			return false
		}
	}
	return false
}
