package climbingtheleaderboard

import (
	"fmt"
	"reflect"
	"testing"
)

func TestClimbingLeaderboard(t *testing.T) {

	result := ClimbingLeaderboard([]int32{100, 100, 50, 40, 40, 20, 10}, []int32{5, 25, 50, 120})
	if reflect.DeepEqual(result, []int32{1, 2, 4, 6}) {
		t.FailNow()
	}
	fmt.Printf("Sample Output 1:\n")
	printResult(result)

	result = ClimbingLeaderboard([]int32{100, 90, 90, 80, 75, 60}, []int32{50, 65, 77, 90, 102})
	if reflect.DeepEqual(result, []int32{1, 2, 4, 5, 6}) {
		t.FailNow()
	}
	fmt.Printf("Sample Output 2:\n")
	printResult(result)
}

func printResult(result []int32) {

	for i, resultItem := range result {
		fmt.Printf("%d", resultItem)

		if i != len(result)-1 {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("\n")
}
