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

func TheBest(nums1 []int, m int, nums2 []int, n int)  {
	cur := m + n - 1
	p1, p2 := m - 1, n - 1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[cur] = nums1[p1]
			p1--
		} else {
			nums1[cur] = nums2[p2]
			p2--
		}

		cur--
	}

	for p2 >= 0 {
		nums1[cur] = nums2[p2]
		p2--
		cur--
	}
}
