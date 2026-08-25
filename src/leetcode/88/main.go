package main

func Merge(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if nums1[m-1] < nums2[n-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[n+m-1] = nums1[m-1]
			m--
		}
	}
}
// Мысли
// Сравниваем ненулевые элементы массивов идя с конца и записываем их в конец
// j это просто m n -=
// индекс для для случая nums1[m]<=nums2[n] только n-= и nums1[m+n]=nums2[n]
// nums1[n]>nums2[m] m-=  nums1[n+m]=nums1[m]

// условие выхода
// n == 0
// пример 2 0    1

// Краевые случаи
// если массив 2 пустой
// 10 1
// m1 n1

