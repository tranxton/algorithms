package algorithms

func Factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * Factorial(n-1)
}

func TailFactorial(n int, result int) int {
	if n == 1 {
		return result
	}

	return TailFactorial(n-1, result*n)
}

func Sum(n []int, i int) int {
	if i == len(n)-1 {
		return n[i]
	}

	return n[i] + Sum(n, i+1)
}

func TailSum(n []int, i int, result int) int {
	if i == len(n)-1 {
		return result + n[i]
	}

	return TailSum(n, i+1, result+n[i])
}

func Count(n []int) int {
	if len(n) == 0 {
		return 0
	}

	return 1 + Count(n[1:])
}

func TailCount(n []int, result int) int {
	if len(n) == 0 {
		return result
	}

	return TailCount(n[1:], result+1)
}

func BiggestNumber(n []int) int {
	if len(n) == 0 {
		return 0
	}

	number := n[0]
	nextNumber := BiggestNumber(n[1:])

	if number > nextNumber {
		return number
	}

	return nextNumber
}

func TailBiggestNumber(n []int, biggestNumber int) int {
	if len(n) == 0 {
		return biggestNumber
	}

	number := n[0]
	if number > biggestNumber {
		return TailBiggestNumber(n[1:], number)
	}

	return TailBiggestNumber(n[1:], biggestNumber)
}
