package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int)  {
	for i := m; i < m + n; i++ {
		nums1[i] = nums2[i -m]
	}
	length := len(nums1)
	// 冒泡算法
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if nums1[i] > nums1[j]{
				temp := nums1[i]
				nums1[i] = nums1[j]
				nums1[j] = temp
			}
		}
	}
}
