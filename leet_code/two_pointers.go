package leet_code

func RemoveElementExcept(nums []int, val int) int {
	k := 0 // Pointer for placement of elements not equal to val

	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}

	return k
}

func HasPair(arr []int, target int) bool {
	seen := make(map[int]bool)

	for _, num := range arr {
		diff := target - num

		if seen[diff] {
			return true
		}

		seen[num] = true
	}

	return false
}

func MoveZerosToEnd(nums []int) []int {
	nonZeroIndex := 0

	for _, num := range nums {
		if num != 0 {
			nums[nonZeroIndex] = num
			nonZeroIndex++
		}
	}

	for i := nonZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}

func TwoSum(nums []int, target int) []int {
	checked := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		if checkedNum, numExist := checked[diff]; numExist {
			return []int{checkedNum, i}
		}

		checked[num] = i
	}

	return []int{}
}

func IsPalindromeTwoPointers(s string) bool {
	firstPointer, secondPointer := 0, len(s)-1

	for firstPointer < secondPointer {
		firstLetter, secondLetter := s[firstPointer], s[secondPointer]

		if !(firstLetter >= 'a' && firstLetter <= 'z') &&
			!(firstLetter >= 'A' && firstLetter <= 'Z') &&
			!(firstLetter >= '0' && firstLetter <= '9') {
			firstPointer++
			continue
		}

		if !(secondLetter >= 'a' && secondLetter <= 'z') &&
			!(secondLetter >= 'A' && secondLetter <= 'Z') &&
			!(secondLetter >= '0' && secondLetter <= '9') {
			secondPointer--
			continue
		}

		if firstLetter >= 'A' && firstLetter <= 'Z' {
			firstLetter += 32
		}
		if secondLetter >= 'A' && secondLetter <= 'Z' {
			secondLetter += 32
		}

		if firstLetter != secondLetter {
			return false
		}

		firstPointer++
		secondPointer--
	}

	return true
}
