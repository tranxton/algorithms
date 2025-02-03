package main

import (
	"algorithms/go_basics"
)

func main() {

}

func passingLeetCodeProblems() {
	/*tree := leet_code.TreeNode{
		Val: 1,
		Left: &leet_code.TreeNode{
			Val: 2,
			Left: &leet_code.TreeNode{
				Val: 4,
				Left: &leet_code.TreeNode{
					Val: 8,
				},
			},
			Right: &leet_code.TreeNode{
				Val: 5,
			},
		},
		Right: &leet_code.TreeNode{
			Val: 3,
			Left: &leet_code.TreeNode{
				Val: 6,
			},
			Right: &leet_code.TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(leet_code.CountNodes(&tree))*/

	/*fmt.Println(leet_code.MissingNumber([]int{1, 2, 4, 5, 6}))
	fmt.Println(leet_code.MySqrt(8))
	fmt.Println(leet_code.Intersection([]int{8, 0, 3}, []int{0, 0}))
	fmt.Println(leet_code.IsPerfectSquare(1664))
	fmt.Println(leet_code.Intersect([]int{2}, []int{4}))
	fmt.Println(leet_code.GuessNumber(10))
	fmt.Println(leet_code.ContainsDuplicate([]int{0, 1, 1}))
	fmt.Println(leet_code.IsValid("}{"))
	fmt.Println(leet_code.IsPalindromeStack(9))*/

	/*nums2 := []int{2, 5, 6}
	nums1 := []int{1, 2, 3, 0, 0, 0}
	leet_code.Merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)*/

	/*pTree := leet_code.TreeNode{
		Val: 1,
		Left: &leet_code.TreeNode{
			Val: 2,
		},
		Right: &leet_code.TreeNode{
			Val: 3,
		},
	}

	qTree := leet_code.TreeNode{
		Val: 1,
		Left: &leet_code.TreeNode{
			Val: 2,
		},
		Right: &leet_code.TreeNode{
			Val: 3,
		},
	}

	fmt.Println(leet_code.IsSameTree(&pTree, &qTree))*/
	/**
		    	   1
		    	 2 | 2
		    3 null | null 3
	 4 null null 4 |
	*/

	//[1,2,2,3,null,null,3,4,null,null,4]
	/*tree := leet_code.TreeNode{
		Val: 1,
		Left: &leet_code.TreeNode{
			Val: 2,
			Left: &leet_code.TreeNode{
				Val: 4,
			},
			Right: &leet_code.TreeNode{
				Val: 5,
			},
		},
		Right: &leet_code.TreeNode{
			Val: 3,
			Right: &leet_code.TreeNode{
				Val: 3,
				Right: &leet_code.TreeNode{
					Val: 4,
				},
			},
			Left: &leet_code.TreeNode{
				Val: 3,
			},
		},
	}*/

	/*fmt.Println(leet_code.IsSymmetricBfs(&tree))
	fmt.Println(leet_code.IsSymmetricDfs(&tree))*/

	/*fmt.Println(leet_code.MaxDepthDfs(&tree))
	fmt.Println(leet_code.MaxDepthBfs(&tree))*/

	//fmt.Println(leet_code.IsBalanced(&tree))

	/*fmt.Println(leet_code.MinDepthDfs(&tree))
	fmt.Println(leet_code.MinDepthBfs(&tree))*/

	/*fmt.Println(leet_code.HasPathSumDfs(&tree, 10))
	fmt.Println(leet_code.HasPathSumBfs(&tree, 10))*/

	/*fmt.Println(leet_code.InvertTreeDfs(&tree))
	fmt.Println(leet_code.InvertTreeBfs(&tree))*/

	/*fmt.Println(leet_code.SumOfLeftLeavesDfs(&tree))
	fmt.Println(leet_code.SumOfLeftLeavesBfs(&tree))*/

	//fmt.Println(leet_code.Search([]int{4, 5, 6, 7, 0, 1, 2}, 2))
}

func testingAlgorithms() {
	/*array := []int{1, 2, 3, 4, 5}
	fmt.Println(algorithms.BinarySearch(array, 1))*/

	/*fmt.Println(algorithms.Factorial(5))
	fmt.Println(algorithms.TailFactorial(5, 1))*/
	/*fmt.Println(algorithms.Sum([]int{1, 2, 3, 4, 5}, 0))
	fmt.Println(algorithms.TailSum([]int{1, 2, 3, 4, 5}, 0, 0))*/

	/*fmt.Println(algorithms.Count([]int{1, 2, 3}))
	fmt.Println(algorithms.TailCount([]int{1, 2, 3}, 0))*/

	/*fmt.Println(algorithms.BiggestNumber([]int{5, 10, 1, 2, 3, 20}))
	fmt.Println(algorithms.TailBiggestNumber([]int{5, 10, 1, 2, 3, 20}, 0))*/

	/*fmt.Println(algorithms.SelectSorting([]int{6, 2, 3, 1, 5, 4}))
	fmt.Println(algorithms.QuickSorting([]int{6, 2, 5, 3, 1, 4, 1, 1, 1, 1, 2}))
	fmt.Println(algorithms.MergeSorting([]int{6, 2, 5, 3, 1, 4, 1, 1, 1, 1, 2}))*/
}

func testingDataStructures() {
	/*phoneBook := data_structures.CreatePhoneBook()

	data_structures.AddNumberToPhoneBook(phoneBook, "Test 1", "111-111-111")
	data_structures.AddNumberToPhoneBook(phoneBook, "Test 2", "222-222-222")
	data_structures.AddNumberToPhoneBook(phoneBook, "Test 3", "333-333-333")

	fmt.Println(phoneBook)
	fmt.Println(data_structures.GetNumberFromPhoneBook(phoneBook, "Test 1"))
	fmt.Println(data_structures.GetNumberFromPhoneBook(phoneBook, "Test 4"))
	fmt.Println(data_structures.GetNameFromPhoneBook(phoneBook, "333-333-333"))
	fmt.Println(data_structures.GetNameFromPhoneBook(phoneBook, "444-444-444"))*/

	/*socialNetwork := data_structures.SocialNetwork{
		Users: []*data_structures.Person{
			{
				Name: "you",
				Friends: []*data_structures.Person{
					{
						Name: "alice",
						Friends: []*data_structures.Person{
							{
								Name: "peggy",
							},
						},
					},
					{
						Name: "bob",
						Friends: []*data_structures.Person{
							{
								Name: "anuj",
							},
							{
								Name: "peggy",
							},
						},
					},
					{
						Name: "clare",
						Friends: []*data_structures.Person{
							{
								Name: "thom",
							},
							{
								Name: "johny",
							},
						},
					},
				},
			},
			{
				Name: "anuj",
			},
			{
				Name: "peggy",
			},
			{
				Name: "jonny",
			},
			{
				Name: "thom",
			},
			{
				Name: "alex",
			},
		},
	}

	fmt.Println(data_structures.BreadthFirstSearch(socialNetwork.Users[0], "thom"))
	fmt.Println(data_structures.DepthFirstSearch(socialNetwork.Users[0], "johny", map[*data_structures.Person]bool{}))
	fmt.Println(data_structures.TailDepthFirstSearch(socialNetwork.Users[0], "you", map[*data_structures.Person]bool{}, []*data_structures.Person{}))*/
}

func testingGoBasics() {
	go_basics.HelloWorld()
	go_basics.FizzBuzz()
	go_basics.TestingErrorProcessing()
	go_basics.DoPanic()
	go_basics.DeferTest()
	go_basics.TestingInteger()
	go_basics.TestingFloatingPointNumbers()
	go_basics.TestingBoolean()
	go_basics.TestingString()
	go_basics.TestingBackticks()
	go_basics.TestingUserTypes()
	go_basics.TestingAliases()
	go_basics.TestingAlternativeWaysOfVariablesDeclaration()
	go_basics.TestingUntypedConstants()
	go_basics.TestingTypedConstants()
	go_basics.TestingEnums()
	go_basics.UserTypesInConstants()
	go_basics.TestingLiterals()
	go_basics.TestingIota()
	go_basics.TestingGlobalScope()
	go_basics.TestingExport()
	go_basics.TestingConditionalOperators()
	go_basics.TestingCycles()
	go_basics.Pointers()
	go_basics.LiteralPointer()
	go_basics.StructPointer()
	go_basics.UserInputCounter()
	go_basics.Arrays()
	go_basics.MultidimensionalArray()
	go_basics.Slices()
	go_basics.ChangingSliceSize()
	go_basics.WorkingWithSlicesInFunc()
	go_basics.ChangingSliceSizeInFunc()
	go_basics.OperationsOnSlices()
	go_basics.TestTaskWithSlices()
}
