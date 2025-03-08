package algorithms

func BinarySearch(array []int, target int) int {
	low, high := 0, len(array)-1

	for low <= high {
		mid := (low + high) / 2

		guess := array[mid]
		if guess == target {
			return mid
		}

		if guess < target {
			low = mid + 1

			continue
		}

		high = mid - 1
	}

	return -1
}
