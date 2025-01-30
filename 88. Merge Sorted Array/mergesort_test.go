package mergesort

import (
	"testing"
)

func TestMergeCase1(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	if !equal(nums1, []int{1, 2, 2, 3, 5, 6}) {
		t.Fail()
	}
}

func TestMergeCase2(t *testing.T) {
	nums1 := []int{1}
	nums2 := []int{0}
	merge(nums1, 1, nums2, 0)
	if !equal(nums1, []int{1}) {
		t.Fail()
	}
}

func TestMergeCase3(t *testing.T) {
	nums1 := []int{0}
	nums2 := []int{1}
	merge(nums1, 0, nums2, 1)
	if !equal(nums1, []int{1}) {
		t.Fail()
	}
}
