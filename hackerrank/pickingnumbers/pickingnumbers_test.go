package pickingnumbers

import (
	"fmt"
	"testing"
)

func TestPickingNumbers(t *testing.T) {

	arr := []int32{4, 6, 5, 3, 3, 1}
	ans := PickingNumbers(arr)
	if ans != 3 {
		t.FailNow()
	}
	fmt.Printf("Sample Output 1: %d\n", ans)

	arr = []int32{1, 2, 2, 3, 1, 2}
	ans = PickingNumbers(arr)
	if ans != 5 {
		t.FailNow()
	}
	fmt.Printf("Sample Output 2: %d\n", ans)

	arr = []int32{1, 99}
	ans = PickingNumbers(arr)
	if ans != 1 {
		t.FailNow()
	}
	fmt.Printf("Sample Output 3: %d\n", ans)

	arr = []int32{1, 50, 50, 51, 98, 99}
	ans = PickingNumbers(arr)
	if ans != 3 {
		t.FailNow()
	}
	fmt.Printf("Sample Output 4: %d\n", ans)
}
