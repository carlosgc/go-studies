package thehurdlerace

import (
	"fmt"
	"testing"
)

func TestHurdleRace(t *testing.T) {

	arr := []int32{1, 6, 3, 5, 2}
	ans := HurdleRace(4, arr)
	if ans != 2 {
		t.FailNow()
	}
	fmt.Printf("Sample Output 1: %d\n", ans)

	arr = []int32{2, 5, 4, 5, 2}
	ans = HurdleRace(7, arr)
	if ans != 0 {
		t.FailNow()
	}
	fmt.Printf("Sample Output 2: %d\n", ans)
}
