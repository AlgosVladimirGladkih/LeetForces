package main

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			name:     "nums1 empty, nums2 has one element",
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
		{
			name:     "nums1 has one element, nums2 has zero",
			nums1:    []int{1, 0},
			m:        1,
			nums2:    []int{0},
			n:        1,
			expected: []int{0, 1},
		},
		{
			name:     "both have elements",
			nums1:    []int{10, 20, 0, 0},
			m:        2,
			nums2:    []int{1, 2},
			n:        2,
			expected: []int{1, 2, 10, 20},
		},
		{
			name:     "both empty",
			nums1:    []int{},
			m:        0,
			nums2:    []int{},
			n:        0,
			expected: []int{},
		},
		{
			name:     "nums1 has elements, nums2 has elements",
			nums1:    []int{5, 6, 7, 0, 0, 0},
			m:        3,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3, 5, 6, 7},
		},
		{
			name:     "nums2 empty",
			nums1:    []int{1, 2, 3},
			m:        3,
			nums2:    []int{},
			n:        0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "overlapping values",
			nums1:    []int{1, 2, 0, 0},
			m:        2,
			nums2:    []int{1, 3},
			n:        2,
			expected: []int{1, 1, 2, 3},
		},
		{
			name:     "nums1 empty, nums2 has two elements",
			nums1:    []int{0, 0},
			m:        0,
			nums2:    []int{1, 2},
			n:        2,
			expected: []int{1, 2},
		},
		{
			name:     "nums1 empty, nums2 has three elements",
			nums1:    []int{0, 0, 0},
			m:        0,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "nums1 has one element, nums2 has two",
			nums1:    []int{1, 0, 0},
			m:        1,
			nums2:    []int{0, 2},
			n:        2,
			expected: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of nums1 to work with
			nums1Copy := make([]int, len(tt.nums1))
			copy(nums1Copy, tt.nums1)
			
			merge(nums1Copy, tt.m, tt.nums2, tt.n)
			
			// Check if result matches expected
			if !equal(nums1Copy, tt.expected) {
				t.Errorf("merge(%v, %d, %v, %d) = %v; expected %v",
					tt.nums1, tt.m, tt.nums2, tt.n, nums1Copy, tt.expected)
			}
		})
	}
}

// Helper function to compare slices
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Example test - fixed to match the function name (merge)
func ExampleMerge() {
	nums1 := []int{1, 3, 5, 0, 0, 0}
	m := 3
	nums2 := []int{2, 4, 6}
	n := 3
	
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
	// Output: [1 2 3 4 5 6]
}