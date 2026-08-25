package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	// Merge from the end
	for n > 0 {
		// If m is 0, we must take from nums2
		// If m > 0, compare and take the larger
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}