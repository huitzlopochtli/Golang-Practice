package main

import (
	"fmt"
)

func ReadYn(sum []uint32, i uint8, low uint8, high uint8) {
	if low == high {
		return
	}

	var yn int8
	fmt.Scan(&yn)

	// Yn (-100 <= Yn <= 100)
	if yn < -100 || yn > 100 {
		panic("Yn (-100 <= Yn <= 100)")
	}

	if yn > 0 {
		sum[i] += uint32(uint16(yn) * uint16(yn))
	}

	ReadYn(sum, i, low+1, high)
}

func ReadTestCases(sum []uint32, index uint8, n_cases uint8) {
	if index == n_cases {
		return
	}

	// X (0 < X <= 100)
	var x uint8
	fmt.Scan(&x)
	if x < 0 || x > 100 {
		panic("X (0 < X <= 100)")
	}

	ReadYn(sum, index, 0, x)

	ReadTestCases(sum, index+1, n_cases)
}

func printArr(arr []uint32, i uint8, n uint8) {
	if i != n {
		fmt.Println(arr[i])
		printArr(arr, i+1, n)
	}
	return
}

func main() {
	// N (1 <= N <= 100)
	var number_of_test_cases uint8
	fmt.Scanln(&number_of_test_cases)
	if number_of_test_cases < 1 || number_of_test_cases > 100 {
		panic("N (1 <= N <= 100")
	}

	// read test cases
	var test_case_sums = make([]uint32, number_of_test_cases)
	ReadTestCases(test_case_sums, 0, number_of_test_cases)

	// print sum
	printArr(test_case_sums, 0, number_of_test_cases)

}

/*
Description

    We want you to calculate the sum of squares of given integers, excluding any negatives.
    The first line of the input will be an integer N (1 <= N <= 100), indicating the number of test cases to follow.
    Each of the test cases will consist of a line with an integer X (0 < X <= 100), followed by another line consisting of X number of space-separated integers Yn (-100 <= Yn <= 100).
    For each test case, calculate the sum of squares of the integers, excluding any negatives, and print the calculated sum in the output.
    Note: There should be no output until all the input has been received.
    Note 2: Do not put blank lines between test cases solutions.
    Note 3: Take input from standard input, and output to standard output.

Rules

    Write your solution using Go Programming Language
    Your source code must be a single file (package main)
    Do not use any for statement
    You may only use standard library packages

*/

/*
Sample Input
2
4
3 -1 1 14
5
9 6 -53 32 16

Sample Output
206
1397
*/
