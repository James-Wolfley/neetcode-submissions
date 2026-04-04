func productExceptSelf(nums []int) []int {
    length := len(nums)
    prefix := make([]int, length)
    prefix[0] = nums[0]
    postfix := make([]int, length)
    postfix[length-1] = nums[length-1]
    for i := 1; i < length; i++ {
        prefix[i] = prefix[i-1] * nums[i]
    }
    for i := length-2; i > 0; i-- {
        postfix[i] = postfix[i+1] * nums[i]
    }
    result := make([]int, length)
    for i := 0; i < length; i++ {
        if i == 0{
            result[i] = postfix[i+1]
        } else if i == length-1{
            result[i] = prefix[i-1]
        } else {
            result[i] = prefix[i-1] * postfix[i+1]
        }
    }
    return result
}
