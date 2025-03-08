package leet_code

import "math"

func IsValid(s string) bool {
	strLen := len(s)
	if strLen%2 != 0 {
		return false
	}

	pList := map[rune]bool{'(': true, '{': true, '[': true, ')': false, '}': false, ']': false}
	var stack []rune

	for _, p := range s {
		if pList[p] {
			stack = append(stack, p)

			continue
		}

		if len(stack) == 0 {
			return false
		}

		stackP := stack[len(stack)-1]

		switch p {
		case ')':
			if stackP != '(' {
				return false
			}

			break
		case '}':
			if stackP != '{' {
				return false
			}

			break
		case ']':
			if stackP != '[' {
				return false
			}

			break
		default:
			return false

		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	xLen := int(math.Log10(float64(x))) + 1
	isOdd := xLen%2 != 0
	xMid := xLen / 2
	var stack []int

	for i := 0; i < xLen; i++ {
		if i < xMid {
			stack = append(stack, x%10)
			x = x / 10

			continue
		}

		if isOdd && i == xMid {
			x = x / 10

			continue
		}

		if stack[len(stack)-1] != x%10 {
			return false
		}

		x = x / 10
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

func twoSum(nums []int, target int) []int {
	for i, num1 := range nums {
		for j, num2 := range nums {
			if num1*num2 == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}
