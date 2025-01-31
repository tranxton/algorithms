package algorithms

func SelectSorting(array []int) []int {
	var sorted []int
	smallestValue := 0

	for range array {
		smallestValueIndex := findSmallestValueIndex(array)
		smallestValue = array[smallestValueIndex]
		array = append(array[:smallestValueIndex], array[smallestValueIndex+1:]...)

		sorted = append(sorted, smallestValue)
	}

	return sorted
}

func findSmallestValueIndex(array []int) int {
	minValue := array[0]
	minValueIndex := 0
	for i := 1; i < len(array); i++ {
		if array[i] < minValue {
			minValue = array[i]
			minValueIndex = i
		}
	}

	return minValueIndex
}

func QuickSorting(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2
	pivot := nums[mid]
	var less, higher []int
	for _, num := range append(nums[mid+1:], nums[:mid]...) {
		if num <= pivot {
			less = append(less, num)

			continue
		}

		higher = append(higher, num)
	}

	return append(append(QuickSorting(less), pivot), QuickSorting(higher)...)
}

func MergeSorting(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2
	left, right := MergeSorting(nums[:mid]), MergeSorting(nums[mid:])
	var merged []int

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]

			continue
		}

		merged = append(merged, right[0])
		right = right[1:]
	}

	return append(append(merged, left...), right...)
}
