package main

import "fmt"

func main() {
	fmt.Println(calculateFib(6))
	fmt.Println(calculateFib(7))
	fmt.Println(calculateFib(8))
	// fmt.Println(calculateFib(50))
	fmt.Println("*****************************")
	memo := make(map[int]int)
	fmt.Println(calculateFibUsingMemo(6, memo))
	memo = nil
	memo = make(map[int]int)
	fmt.Println(calculateFibUsingMemo(7, memo))
	memo = nil
	memo = make(map[int]int)
	fmt.Println(calculateFibUsingMemo(8, memo))
	memo = nil
	memo = make(map[int]int)
	fmt.Println(calculateFibUsingMemo(50, memo))

}
