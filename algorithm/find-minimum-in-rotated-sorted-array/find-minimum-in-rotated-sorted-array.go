package find_minimum_in_rotated_sorted_array

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[1]
	}
	l, r := 0, len(nums)-1
	if nums[l] <= nums[r] {
		return nums[l]
	}
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		if nums[0] < nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

func findMin2(nums []int) int {
	n := len(nums)
	l, r := 0, n-1

	for l < r {
		mid := l + r + 1>>1
		if nums[mid] >= nums[0] {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if r+1 < n {
		return nums[r+1]
	} else {
		return nums[0]
	}
}
