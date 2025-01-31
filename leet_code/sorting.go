package leet_code

func Merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1, quickSort3(append(nums1[:m], nums2[:n]...)))
}

func quickSort3(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	var less, higher []int
	mid := len(nums) / 2
	pivot := nums[mid]

	for _, num := range append(nums[:mid], nums[mid+1:]...) {
		if num <= pivot {
			less = append(less, num)

			continue
		}

		higher = append(higher, num)
	}

	return append(append(quickSort3(less), pivot), quickSort3(higher)...)
}

func ContainsDuplicate(nums []int) bool {
	numsLen := len(nums)
	duplicates := make(map[int]bool, numsLen)
	secondI := numsLen

	for firstI := 0; firstI < numsLen && firstI < secondI; firstI++ {
		if duplicates[nums[firstI]] {
			return true
		}

		duplicates[nums[firstI]] = true

		secondI = numsLen - firstI - 1
		if firstI == secondI {
			return false
		}

		if duplicates[nums[secondI]] {
			return true
		}

		duplicates[nums[secondI]] = true
	}

	return false
}

func quickSort4(nums []int) ([]int, bool) {
	var duplicates map[int]bool

	var less, higher []int
	mid := len(nums) / 2
	pivot := nums[mid]

	duplicates[pivot] = true

	for _, num := range append(nums[:mid], nums[mid+1:]...) {
		if duplicates[num] {
			return make([]int, 0), true
		}
		duplicates[num] = true

		if num <= pivot {
			less = append(less, num)

			continue
		}

		higher = append(higher, num)
	}

	sortedLess, foundDuplicateL := quickSort4(less)
	sortedHigher, foundDuplicateH := quickSort4(higher)

	if foundDuplicateL || foundDuplicateH {
		return make([]int, 0), true
	}

	return append(append(sortedLess, pivot), sortedHigher...), false
}
