package sort

// 归并排序

func MergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	middle := n / 2
	left := nums[:middle]
	right := nums[middle:]

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left []int, right []int) []int {
	result := make([]int, 0)

	for len(left)>0 && len(right)>0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		result = append(result, left...)
	}
	if len(right) > 0 {
		result = append(result, right...)
	}

	return result
}
