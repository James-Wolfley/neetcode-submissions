func topKFrequent(nums []int, k int) []int {
    frequency := make(map[int]int)
    for _, v := range nums {
        frequency[v]++
    }
    sorted := make([][]int, len(nums)+1)
    for key, freq := range frequency {
        sorted[freq] = append(sorted[freq], key)
    }

    var result []int
    for i := len(nums); i >= 0; i--{
        for _, v := range sorted[i] {
            result = append(result, v)
            if len(result) == k {
                return result
            }
        }
    }
    return result
}
