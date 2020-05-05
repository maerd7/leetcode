package main

//给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//
// 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
// 你可以假设 nums1 和 nums2 不会同时为空。
//
// 示例 1:
//
// nums1 = [1, 3]
//nums2 = [2]
//
//则中位数是 2.0
//
//
// 示例 2:
//
// nums1 = [1, 2]
//nums2 = [3, 4]
//
//则中位数是 (2 + 3)/2 = 2.5
//
// Related Topics 数组 二分查找 分治算法

// [1,2,3,4,5] [6]   -- [3.5]
// [1] [2]   1.5

// [1,2,3,4,5,6] [7,8,9,10]  5.5

// [1,2,3] [1,2,3]  --> 2

//leetcode submit region begin(Prohibit modification and deletion)
// 最差情况是 m + n 因此算法复杂度为 O(M+N) 慢于 log(m+n)
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	m := 0
	n := 0
	num := make([]int, len1+len2)
	for i := 0; i < len1+len2; i++ {

		if m == len1 {
			num[i] = nums2[n]
			n++
		} else if n == len2 {
			num[i] = nums1[m]
			m++
		} else if nums1[m] > nums2[n] {
			num[i] = nums2[n]
			n++
		} else {
			num[i] = nums1[m]
			m++
		}
	}
	if len(num)%2 == 0 {
		return float64((num[len(num)/2] + num[len(num)/2-1]) / 2)
	} else {
		return float64(num[len(num)/2])
	}
}

//leetcode submit region end(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	mid1 := (len1 + 1) / 2
	mid2 := (len2 + 1) / 2
	// 递归出口条件

	// 分情况讨论，按第一位数值比大小
	// nums1 --> a,b
	// nums2 --> c,d
	// abcd acbd acdb
	// cdab cabd cdba
	a := nums1[0]
	b := nums1[mid1]
	c := nums2[0]
	d := nums2[mid2]
	arrayA := nums1[0:mid1]
	arrayB := nums1[mid1:]
	arrayC := nums2[0:mid2]
	arrayD := nums2[mid2:]
	// a <= c
	if a <= c {
		// b <= c abcd
		if b <= c {
			return findMedianSortedArrays(arrayB, arrayC)
			// b > c
		} else {
			// b <= d acbd
			if b <= d {
				return findMedianSortedArrays(arrayC, arrayB)
				// b > d acdb
			} else {
				return findMedianSortedArrays(arrayC, arrayD)
			}
		}
		// a > c
	} else {
		// a <= d
		if a <= d {
			// b <= d cabd
			if b <= d {
				return findMedianSortedArrays(arrayA, arrayB)
				// b > d cadb
			} else {
				return findMedianSortedArrays(arrayA, arrayD)
			}
			// a > d cdab
		} else {
			return findMedianSortedArrays(arrayD, arrayA)
		}
	}
}
