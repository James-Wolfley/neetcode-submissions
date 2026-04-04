func longestConsecutive(nums []int) int {
    numSet := make(map[int]struct{})
    for _, v := range nums {
        numSet[v] = struct{}{}
    }

    longest := 0
    for v := range numSet{
        if _, found := numSet[v-1]; !found{
            length := 1
            for {
                if _, exists := numSet[v+length]; exists {
                    length++
                } else {
                    break
                }
            }
            if length > longest {
                longest = length
            }
        }
    }
    return longest
}
