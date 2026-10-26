package main

import "fmt"

// Leetcode function
func LongestSubstring(nums []int) int {
	var count, save_count, rez int
	for _, elem := range nums {
		count++
		if elem == 0 {
			if save_count+count-1 > rez {
				rez = save_count + count - 1
			}
			save_count = count - 1
			count = 0
			// if save_count-1 > rez {
			// rez = save_count - 1
			// }
		} else if save_count+count > rez {
			rez = save_count + count
		}
	}
	if save_count == 0 {
		rez = count - 1
	}
	if rez < 0 {
		rez = 0
	}
	return max(rez, count+save_count-1)
}

func main() {
	nums1 := []int{1, 1, 1, 1}
	fmt.Printf("%d\n", LongestSubstring(nums1))
}
