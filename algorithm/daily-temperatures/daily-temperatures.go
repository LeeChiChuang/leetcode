package daily_temperatures
	
func dailyTemperatures(T []int) []int {
	n := len(T)
	if n == 0 {
		return T
	}

	st := make([]int, 0)
	ans := make([]int, n)
	for i:=n-1; i>=0; i-- {
		for len(st) != 0 && T[i] >= T[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			ans[i] = 0
		} else {
			ans[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}

	return ans
}
