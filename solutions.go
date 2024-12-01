package main

import (
	"fmt"
	"sort"
	"strconv"
)

// MiniMaxSum solution
func miniMaxSum() {
	arr := make([]int64, 5)
	fmt.Println("Enter 5 space-separated integers:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}

	// Sort the array
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var minSum, maxSum int64
	for i := 0; i < 4; i++ {
		minSum += arr[i]
	}
	for i := 1; i < 5; i++ {
		maxSum += arr[i]
	}

	fmt.Printf("Result: %d %d\n", minSum, maxSum)
}

// Plus Minus solutions
func plusMinus() {
	// Get the size of array
	var n int
	fmt.Println("Enter the size of array:")
	fmt.Scan(&n)

	// Input validation
	if n <= 0 || n > 100 {
		fmt.Println("Array size must be between 1 and 100")
		return
	}

	// Get the array elements
	arr := make([]int, n)
	fmt.Printf("Enter %d space-separated integers:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		// Input validation
		if arr[i] < -100 || arr[i] > 100 {
			fmt.Println("Array elements must be between -100 and 100")
			return
		}
	}

	// Count positive, negative, and zero values
	var positiveCount, negativeCount, zeroCount float64

	for _, num := range arr {
		switch {
		case num > 0:
			positiveCount++
		case num < 0:
			negativeCount++
		default:
			zeroCount++
		}
	}

	// Calculate proportions
	n64 := float64(n)
	positiveProportion := positiveCount / n64
	negativeProportion := negativeCount / n64
	zeroProportion := zeroCount / n64

	// Print results with 6 decimal places
	fmt.Printf("%.6f\n", positiveProportion)
	fmt.Printf("%.6f\n", negativeProportion)
	fmt.Printf("%.6f\n", zeroProportion)
}

// Time Conversion solutions
func timeConversion() {
	var timeStr string
	fmt.Println("Enter time in 12-hour format (hh:mm:ssAM/PM):")
	fmt.Scan(&timeStr)

	// Extract hours, minutes, seconds, and period
	hours := timeStr[0:2]
	minutes := timeStr[3:5]
	seconds := timeStr[6:8]
	period := timeStr[8:10]

	var convertedTime string

	// Convert hours based on period (AM/PM)
	switch period {
	case "PM":
		if hours != "12" {
			// Convert hours to integer, add 12, convert back to string
			h, _ := strconv.Atoi(hours)
			convertedTime = fmt.Sprintf("%02d", h+12)
		} else {
			convertedTime = "12" // 12 PM stays as 12
		}
	case "AM":
		if hours == "12" {
			convertedTime = "00" // 12 AM becomes 00
		} else {
			convertedTime = hours // Other AM hours stay the same
		}
	}

	// Format and print the result
	militaryTime := fmt.Sprintf("%s:%s:%s", convertedTime, minutes, seconds)
	fmt.Printf("converted time: %s\n", militaryTime)
}
