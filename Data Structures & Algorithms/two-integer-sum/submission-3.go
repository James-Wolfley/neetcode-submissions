func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    for i, v := range nums{
        val, ok := seen[target - v]
        if ok {
            return []int{val, i}
        }else{
            seen[v] = i
        }
    }
    return []int{}
}
