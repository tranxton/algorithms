package main

func main() {
	testingAlgorithms()
	testingDataStructures()
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
