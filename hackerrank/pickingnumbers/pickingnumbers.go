// Package pickingnumbers provides a solution implementation for the Picking Numbers problem.
//
// Statement: Given an array of integers, find the longest subarray where the absolute difference between any two elements is less than or equal to 1.
//
// A complete description of this problem can be found here: https://www.hackerrank.com/challenges/picking-numbers
package pickingnumbers

// PickingNumbers finds the longest subarray in a given array of integers where the absolute difference between any two elements
// is less than or equal to 1.
func PickingNumbers(arr []int32) int32 {

	// Creating the numberQtd array to store the quantity for each number in array arr
	var numberQtd [101]int32

	maxSubArrLen := int32(0)
	n := len(arr)
	for idArr := 0; idArr < n; idArr++ {

		idNum := arr[idArr]
		numberQtd[idNum] = numberQtd[idNum] + 1
		nQtd := numberQtd[idNum]

		// The longest subarray length are given by the max sum of the two consecutive elements in numberQtd  (i.e., numberQtd[i] and numberQtd[i+1]).

		// Calculating the length of the subarray that contain the numbers (numberQtd[idNum-1], numberQtd[idNum])
		subArrLen := numberQtd[idNum-1] + nQtd
		if subArrLen > maxSubArrLen {
			maxSubArrLen = subArrLen
		}

		// Calculating the length of the subarray that contain the numbers (numberQtd[idNum], numberQtd[idNum+1])
		subArrLen = numberQtd[idNum+1] + nQtd
		if subArrLen > maxSubArrLen {
			maxSubArrLen = subArrLen
		}
	}

	return maxSubArrLen
}
