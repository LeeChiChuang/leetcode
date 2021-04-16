package next_greater_element_i

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0)
	for i:=0; i<len(nums1); i++ {
		minVal := -1
		//找位置
		k := 0
		for nums2[k] != nums1[i] {
			k++
		}
		for j:=k+1; j<len(nums2); j++ {
			if nums2[j] > nums1[i] {
				minVal = nums2[j]
				break
			}
		}
		ans = append(ans, minVal)
	}

	return ans
}

// 单调栈
// 时间复杂度O(n) 每个元素只会入栈出栈一次
func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	n2 := len(nums2)
	st := make([]int, n2)
	m := make(map[int]int)
	for i:=n2-1; i>=0; i-- {
		for len(st) != 0 && nums2[i] > st[len(st)-1] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			m[nums2[i]] = -1
		} else {
			m[nums2[i]] = st[len(st)-1]
		}
		st = append(st, nums2[i])
	}

	ans := make([]int, len(nums1))
	for i, v := range nums1 {
		ans[i] = m[v]
	}

	return ans
}