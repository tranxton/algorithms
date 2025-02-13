package leet_code

import (
	"algorithms/algorithms"
	"fmt"
	"strconv"
	"strings"
)

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

func MaxDistanceToNearestPerson(seats []int) int {
	prevSeat, maxSeat := -1, 0

	for i, seat := range seats {
		if seat == 1 {
			if prevSeat == -1 {
				maxSeat = i
			} else {
				maxSeat = max(maxSeat, (i-prevSeat)/2)
			}

			prevSeat = i
		}
	}

	maxSeat = max(maxSeat, len(seats)-1-prevSeat)

	return maxSeat
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func LongestMonotonicSubarray(arr []int) []int {
	if len(arr) < 2 {
		return []int{0, 0}
	}

	direction, start, maxStart, maxEnd := 0, 0, 0, 0

	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]

		if diff > 0 {
			if direction == -1 {
				start = i - 1
			}

			direction = 1
		} else if diff < 0 {
			if direction == 1 {
				start = i - 1
			}

			direction = -1
		} else {
			direction = 0
			start = i
		}

		if i-start > maxEnd-maxStart {
			maxStart, maxEnd = start, i
		}
	}

	return []int{maxStart, maxEnd}
}

func MakeRanges(nums []int) string {
	if len(nums) == 0 {
		return ""
	}

	sortedNums := algorithms.QuickSorting(nums)
	start, end := sortedNums[0], sortedNums[0]
	var result []string

	for i := 1; i < len(sortedNums); i++ {
		if end == sortedNums[i]-1 {
			end = sortedNums[i]

			continue
		}

		if start == end {
			result = append(result, strconv.Itoa(end))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", start, end))
		}

		start, end = sortedNums[i], sortedNums[i]
	}

	if start == end {
		result = append(result, strconv.Itoa(end))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", start, end))
	}

	return strings.Join(result, ",")
}
