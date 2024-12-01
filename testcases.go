package main

import "fmt"

type TestCase struct {
	name     string
	function func()
}

func RunTestCases() {
	testCases := []TestCase{
		{name: "1. Mini-Max Sum", function: miniMaxSum},
		{name: "2. Plus Minus", function: plusMinus},
		{name: "3. Time Conversion", function: timeConversion},
	}

	for {
		displayMenu(testCases)
		choice := getUserChoice()

		if choice == 0 {
			fmt.Println("Exiting...")
			break
		}

		runSelectedTest(choice, testCases)
	}
}

func displayMenu(testCases []TestCase) {
	fmt.Println("\n=== Problem Solving Test Cases ===")
	for _, tc := range testCases {
		fmt.Println(tc.name)
	}
	fmt.Println("0. Exit")
	fmt.Println("===============================")
}

func getUserChoice() int {
	var choice int
	fmt.Print("Choose a problem (0 to exit): ")
	fmt.Scan(&choice)
	return choice
}

func runSelectedTest(choice int, testCases []TestCase) {
	if choice > 0 && choice <= len(testCases) {
		fmt.Printf("\nRunning: %s\n", testCases[choice-1].name)
		testCases[choice-1].function()
	} else {
		fmt.Println("Invalid choice! Please try again.")
	}
}
