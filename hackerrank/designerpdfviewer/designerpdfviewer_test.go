package designerpdfviewer

import (
	"fmt"
	"testing"
)

func TestDesignerPdfViewer(t *testing.T) {

	arr := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	ans := DesignerPdfViewer(arr, "abc")
	if ans != 9 {
		t.FailNow()
	}
	fmt.Printf("Sample Output 1: %d\n", ans)

	arr = []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}
	ans = DesignerPdfViewer(arr, "zaba")
	if ans != 28 {
		t.FailNow()
	}
	fmt.Printf("Sample Output 2: %d\n", ans)
}
