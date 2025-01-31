package leet_code

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SearchInsert(nums []int, target int) int {
	low, guess, targetInsertIndex, high := 0, 0, 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		guess = nums[mid]

		if guess == target {
			return mid
		}

		if guess < target {
			low = mid + 1
			targetInsertIndex = mid + 1

			continue
		}

		high = mid - 1
		targetInsertIndex = mid - 1
	}

	if guess > target {
		return targetInsertIndex + 1
	}

	if targetInsertIndex >= 0 {
		return targetInsertIndex
	}

	return 0
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		guess := nums[mid]

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

func GetCommon(nums1 []int, nums2 []int) int {
	for _, num1 := range nums1 {
		low, high := 0, len(nums2)-1

		for low <= high {
			mid := (low + high) / 2
			guess := nums2[mid]

			if guess == num1 {
				return guess
			}

			if guess < num1 {
				low = mid + 1

				continue
			}

			high = mid - 1
		}
	}

	return -1
}

func MissingNumber(nums []int) int {
	vSum, iSum := getNumbersSum(nums, 0, 0, 0)

	return iSum - vSum
}

func getNumbersSum(nums []int, i int, vResult int, iResult int) (int, int) {
	if i == len(nums)-1 {
		return vResult + nums[i], iResult + (i + 1)
	}

	return getNumbersSum(nums, i+1, vResult+nums[i], iResult+(i+1))
}

func getDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	return getDepth(root.Left, depth+1)
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth, rightDepth := getDepth(root.Left, 0), getDepth(root.Right, 0)

	fmt.Println(root.Val, leftDepth, rightDepth, 1<<leftDepth, 1<<rightDepth)

	if leftDepth == rightDepth {
		return (1 << leftDepth) + countNodes(root.Right)
	}

	return (1 << rightDepth) + countNodes(root.Left)
}

func CountNodes(root *TreeNode) int {
	return countNodes(root)
}

func FirstBadVersion(n int) int {
	return firstBadVersion(n)
}

func firstBadVersion(n int) int {
	left, right := 1, n

	for left <= right {
		mid := (left + right) / 2

		if isBadVersion(mid) {
			right = mid - 1

			continue
		}

		left = mid + 1
	}

	return left
}

func isBadVersion(n int) bool {
	return n >= 556
}

func MySqrt(x int) int {
	l, m, h := 1, 1, x

	for l <= h {
		m = (l + h) / 2

		pow := m * m
		fmt.Println(l, h, m, pow, x)

		if pow == x {
			return m
		}

		if pow < x {
			l = m + 1

			continue
		}

		h = m - 1
	}

	return h
}

func Intersection(nums1 []int, nums2 []int) []int {
	return intersection(nums1, nums2)
}

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums2) > len(nums1) {
		nums1, nums2 = nums2, nums1
	}

	var intersections []int
	sortedNums1, sortedNums2, prevNum2 := quickSort(nums1), quickSort(nums2), -1

	for _, num2 := range sortedNums2 {
		if prevNum2 == num2 {
			continue
		}

		l, h := 0, len(sortedNums1)

		for l <= h {
			mid := (l + h) / 2
			if mid >= len(sortedNums1) {
				break
			}

			num1 := sortedNums1[mid]

			if num1 == num2 {
				prevNum2 = num2
				intersections = append(intersections, num2)
				sortedNums1 = append(sortedNums1[:mid], sortedNums1[mid+1:]...)

				break
			}

			if num1 < num2 {
				l = mid + 1

				continue
			}

			h = mid - 1
		}
	}

	return intersections
}

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2
	pivot := nums[mid]
	var less, greater []int

	for _, num := range append(nums[:mid], nums[mid+1:]...) {
		if num <= pivot {
			less = append(less, num)

			continue
		}

		greater = append(greater, num)
	}

	return append(append(quickSort(less), pivot), quickSort(greater)...)
}

func IsPerfectSquare(num int) bool {
	return isPerfectSquare(num)
}

func isPerfectSquare(num int) bool {
	l, h := 0, num

	for l <= h {
		mid := (l + h) / 2
		sqrt := mid * mid

		if sqrt == num {
			return true
		}

		if sqrt < num {
			l = mid + 1

			continue
		}

		h = mid - 1
	}

	return false
}

func Intersect(nums1 []int, nums2 []int) []int {
	return intersect(nums1, nums2)
}

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums2) > len(nums1) {
		nums1, nums2 = nums2, nums1
	}

	var intersects []int
	sortedNums1 := quickSort2(nums1)
	for _, num2 := range nums2 {
		l, h := 0, len(sortedNums1)

		for l <= h {
			mid := (l + h) / 2
			if mid >= len(sortedNums1) {
				break
			}

			num1 := sortedNums1[mid]

			if num1 == num2 {
				intersects = append(intersects, num2)
				sortedNums1 = append(sortedNums1[:mid], sortedNums1[mid+1:]...)

				break
			}

			if num1 < num2 {
				l = mid + 1

				continue
			}

			h = mid - 1
		}
	}

	return intersects
}

func quickSort2(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	var less, greater []int
	mid := len(nums) / 2
	pivot := nums[mid]

	for _, num := range append(nums[:mid], nums[mid+1:]...) {
		if num <= pivot {
			less = append(less, num)

			continue
		}

		greater = append(greater, num)
	}

	return append(append(quickSort2(less), pivot), quickSort2(greater)...)
}

func GuessNumber(n int) int {
	return guessNumber(n)
}

func guessNumber(n int) int {
	l, h := 1, n

	for l <= h {
		mid := (l + h) / 2
		switch guessNum(mid) {
		case 1:
			l = mid + 1
		case -1:
			h = mid - 1
		case 0:
			return mid
		}
	}

	return -1
}

func guessNum(num int) int {
	guessed := 6
	switch {
	case num > guessed:
		return -1
	case num < guessed:
		return 1
	default:
		return 0
	}
}

func Search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[low] <= nums[mid] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}

		}
	}

	return -1
}
