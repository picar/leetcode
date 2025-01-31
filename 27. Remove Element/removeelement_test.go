package removeelement

import (
	"testing"
)

func TestRemoveElementCase1(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	if RemoveElement(nums, 3) != 2 {
		t.Fail()
	}
}

func TestRemoveElementCase2(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	if RemoveElement(nums, 2) != 5 {
		t.Fail()
	}
}
