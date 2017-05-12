package main

func contains(slice []int64, element int64) bool {
	for _, current := range slice {
		if element == current {
			return true
		}
	}
	return false
}
