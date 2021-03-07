package algorithm

func Partition(s string) [][]string {
	var res [][]string          // 待返回的解集
	dfs([]string{}, 0, &res, s) // 当前的部分解为空数组，从索引0开始，往后切回文串
	return res
}

// dfs：基于当前的部分解temp，去切从start到末尾的子串
func dfs(temp []string, start int, res *[][]string, s string) {
	if start == len(s) { // 当start指针越界了，一直切出回文才走到这，将当前的部分解temp加入解集res
		t := make([]string, len(temp)) // 新建一个和temp等长的切片
		copy(t, temp)                  // temp还要在递归中继续被修改，不能将它的引用推入res
		*res = append(*res, t)         // 将temp的拷贝 加入解集res
		return                         // 结束掉当前这个递归（分支）
	}

	for i := start; i < len(s); i++ { // 枚举出当前的所有选项，从索引start到末尾索引
		if isPali(s, start, i) { // 当前选择i，如果 start到 i 是回文串，就切它
			temp = append(temp, s[start:i+1]) // 切出来，加入到部分解temp
			dfs(temp, i+1, res, s)            // 基于这个选择，继续往下递归，继续切
			temp = temp[:len(temp)-1]         // 上面递归结束了，撤销当前选择i，去下一轮迭代
		}
	}
}
// 双指针 判断字符串s是否回文
func isPali(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}