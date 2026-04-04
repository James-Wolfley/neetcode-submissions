func maxSlidingWindow(nums []int, k int) []int {
    result := []int{}
	q := []int{}
	l := 0
	r := 0

	for r < len(nums) {
		for len(q) > 0 && nums[q[len(q)-1]] < nums[r] {
			q = q[:len(q)-1]
		}
		q = append(q, r)

		if l > q[0] {
			q = q[1:]
		}

		if (r + 1) >= k {
			result = append(result, nums[q[0]])
			l += 1
		}
		r += 1
	}
	
	return result
}
