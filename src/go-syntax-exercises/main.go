package main

import(
    "fmt"
	runpkg "Go-syntax-exercises/runningSum"
) 



func main(){
	exampleArr := [5]int {1,2,3,4,5}
	sumArr := runpkg.RunningSumfunc(exampleArr[:])
	fmt.Println(sumArr)
}
