package linearsearch

func linear_search(arr []int, target int) bool {
	for _, num := range arr {
		if num == target {
			return true
		}
	}

	return false
}
