func maxArea(heights []int) int {
	l := 0
	r := len(heights)-1
	max := 0
	for l < r {
		var vol int
		if heights[l] < heights[r]{
			vol = heights[l]
		} else {
			vol = heights[r]
		}
		vol *= r-l
		if vol > max {
			max = vol
		}
		if heights[l] < heights[r] {
			l++
		}else{
			r--
		}
	}
	return max
}
