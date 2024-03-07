package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	numberList := generateNumList(100000000)
	middle := len(numberList) / 2
	firstHalf := numberList[:middle]
	secondHalf := numberList[middle:]

	var wg sync.WaitGroup
	var sum1, sum2 int

	wg.Add(2)
	go func() {
		defer wg.Done()
		sum1 = sumList(firstHalf)
	}()
	go func() {
		defer wg.Done()
		sum2 = sumList(secondHalf)
	}()
	wg.Wait()

	total := sum1 + sum2
	fmt.Println("Total:", total)
}

func sumList(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total 
}

func generateNumList(length int) []int {
	nums := make([]int, length)
	for i := 0; i < length; i++ {
		nums[i] = randomNumberUnder(100)
	}
	return nums
}

func randomNumberUnder(num int) int {
	return rand.Int() % num
}