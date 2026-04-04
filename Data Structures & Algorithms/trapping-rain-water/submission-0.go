func trap(height []int) int {
	n := len(height)
	premax := make([]int, n)
	posmax := make([]int, n)

	premax[0] = height[0]
	posmax[n-1] = height[n-1]
	for i := 1; i < n; i++ {
    	premax[i] = max(premax[i-1], height[i])
    	posmax[n-1-i] = max(posmax[n-i], height[n-1-i])
	}
	var total int
	for i := 0; i < n; i++ {
		total +=  min(premax[i], posmax[i]) - height[i]
	}
	return total
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
