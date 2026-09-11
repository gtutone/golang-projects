package main

import (
	"fmt"
	"testing"
)

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	risingCost := 0.0

	for counter := 0; counter < numMessages; counter++ {
		risingCost = float64(counter) * 0.01
		totalCost = totalCost + 1 + risingCost
	}

	return totalCost // return the total cost (as a float64) to send a batch of numMessages messages.
}

// 1st message: 1.0 + 0.00
// 2nd message: 1.0 + 0.01
// 3rd message: 1.0 + 0.02
// 4th message: 1.0 + 0.03

// ############################
// # Unit tests go under here #
// ############################

func Test1(t *testing.T) {
	type testCase struct {
		numMessages int
		expected    float64
	}
	runCases := []testCase{
		{10, 10.45},
		{20, 21.9},
	}

	submitCases := append(runCases, []testCase{
		{0, 0.0},
		{1, 1.0},
		{5, 5.10},
		{30, 34.35},
	}...)

	testCases := runCases
	// if withSubmit {
	// 	testCases = submitCases
	// }
	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := bulkSend(test.numMessages)
		if fmt.Sprintf("%.2f", output) != fmt.Sprintf("%.2f", test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %.2f
Actual:     %.2f
Fail
`, test.numMessages, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %.2f
Actual:     %.2f
Pass
`, test.numMessages, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
// var withSubmit = true
