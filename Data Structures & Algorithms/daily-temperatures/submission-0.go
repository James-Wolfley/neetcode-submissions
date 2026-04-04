func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0)

	for i, v := range temperatures {
		for len(stack) > 0 && v > temperatures[stack[len(stack)-1]]{
			elem := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[elem] = i-elem
		}
		stack = append(stack, i)
	}
	return res
}
