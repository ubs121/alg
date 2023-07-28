package common

func FindIntersection(set1, set2 map[int]bool) map[int]bool {
	intersection := make(map[int]bool)

	// Check each element in set1 if it exists in set2
	for k := range set1 {
		if _, exists := set2[k]; exists {
			intersection[k] = true
		}
	}

	return intersection
}
