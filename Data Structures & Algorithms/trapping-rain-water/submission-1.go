func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	l := 0
	r := n-1

	leftMax := height[l]
	rightMax := height[r]
	result := 0

	for l < r {
		if leftMax < rightMax {
			l += 1
			leftMax = max(leftMax, height[l])
			result += leftMax - height[l]
		} else {
			r -= 1
			rightMax = max(rightMax, height[r])
			result += rightMax - height[r]
		}
	}

	return result
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