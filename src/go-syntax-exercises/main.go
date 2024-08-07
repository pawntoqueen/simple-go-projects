package main

import "fmt"

func main(){
	exampleArr := [5]int {1,2,3,4,5}
	sumArr := runningSum(exampleArr[:])
	fmt.Println(sumArr)
}


func runningSum(nums []int) []int {
    sum := 0
    retArr := make([]int, len(nums))
    
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        retArr[i] = sum
    }
    return retArr
}