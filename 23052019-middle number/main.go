package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	//找最大值
	if m > n {
		tmp := nums1
		nums1 = nums2
		nums2 = tmp
		t := m
		m = n
		n = t
	}
	//中位数= (n+m)/2
	//定位i 在m中， n中 (n+m)/2-i
	iMax := m
	iMin := 0
	mid := (m + n + 1) / 2
	for {

		if iMin > iMax {
			break
		}
		i := (iMax + iMin) / 2
		j := mid - i
		if i < iMax && nums1[i] < nums2[j-1] {
			iMin = i + 1
		} else if i > iMin && nums1[i] > nums2[j-1] {
			iMax = i - 1
		} else {
			left := 0.0
			if i == 0 {
				left = float64(nums2[j-1])
			} else if j == 0 {
				left = float64(nums1[i-1])
			} else {
				left = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			}
			if (m+n)%2 == 1 {
				return left
			}
			right := 0.0
			if i == m {
				right = float64(nums2[j])
			} else if j == n {
				right = float64(nums1[i])
			} else {
				right = math.Min(float64(nums1[i]), float64(nums2[j]))
			}
			return (right + left) / 2.0
		}

	}
	return 0.0
}
func main() {

}
