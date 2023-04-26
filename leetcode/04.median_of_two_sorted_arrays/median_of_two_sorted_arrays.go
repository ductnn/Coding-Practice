package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//
	m, n := len(nums1), len(nums2)
	mid := (m + n + 1) >> 1

	l, r := 0, m
	i, j := 0, 0

	for l <= r {
		i := (l + r) >> 1
		j := mid - i

		if i > 0 && nums1[i-1] > nums2[j] {
			r = i - 1
		} else if i != len(nums1) && nums1[i] < nums2[j-1] {
			l = i + 1
		} else {
			break
		}
	}

	midLeft, midRight := 0, 0
	if i == 0 {
		midLeft = nums2[j-1]
	} else if j == 0 {
		midLeft = nums1[i-1]
	} else {
		midLeft = max(nums1[i-1], nums2[j-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if i == len(nums1) {
		midRight = nums2[j]
	} else if j == len(nums2) {
		midRight = nums1[i]
	} else {
		midRight = min(nums1[i], nums2[j])
	}
	return float64(midLeft+midRight) / 2
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	//
}
