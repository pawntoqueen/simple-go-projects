package runningSum


func RunningSumfunc(nums []int) []int {
    sum := 0
    retArr := make([]int, len(nums))
    
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        retArr[i] = sum
    }
    return retArr
}
