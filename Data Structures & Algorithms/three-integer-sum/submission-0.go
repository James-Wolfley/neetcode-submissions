func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i, val := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums)-1
		for l < r {
			if nums[l]+nums[r] < -val{
				l++
			}else if nums[l]+nums[r] > -val{
				r--
			} else {
				result = append(result, []int{nums[i],nums[l],nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return result
}
