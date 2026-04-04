func longestConsecutive(nums []int) int {
    if len(nums) == 0{
        return 0
    }
    seq := make(map[int]bool)
    for _, v := range nums {
        seq[v] = true
    }

    longest := 1
    for _, v := range nums{
        if !seq[v-1] {
            continue
        }
        i := 1
        count := 1
        for {
            count++
            if !seq[v+i] {
                if count > longest{
                    longest = count
                }
                break
            }
            i++
        }
    }
    return longest
}
