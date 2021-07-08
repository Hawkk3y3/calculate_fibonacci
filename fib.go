package main

func calculateFib(number int) int {
	//calculates fibonachi number at number index provided
	if number <= 2 {
		return 1
	}
	return calculateFib(number-1) + calculateFib(number-2)
}

func calculateFibUsingMemo(number int, memo map[int]int) int {
	//calculates fibonachi number at number index provided with Memoization
	if val, ok := memo[number]; ok {
		return val
	}
	if number <= 2 {
		return 1
	}
	memo[number] = calculateFibUsingMemo(number-1, memo) + calculateFibUsingMemo(number-2, memo)
	return memo[number]
}
