func twoSum(numbers []int, target int) []int {
    existing := make(map[int]int)
    for key, val := range numbers {
        if v, okay := existing[target-val]; okay{
            return []int{v+1, key+1}
        }
        existing[val] = key
    }
    return []int{0,1}
}
