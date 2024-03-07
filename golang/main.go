package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	numberList := generateNumList(100000000)
	numberOfGroups := 5
	splitList := splitList(numberList, numberOfGroups)

	ch := make(chan int64, numberOfGroups)
	var wg sync.WaitGroup

	for _, list := range splitList {
		wg.Add(1)
		go beginSumming(list, ch, &wg)
	}
	wg.Wait()

	var totalSum int64
	close(ch)
	for num := range ch {
		totalSum += num
	}

	fmt.Println("Total sum:", totalSum)
}

func beginSumming(nums []int64, ch chan int64, wg *sync.WaitGroup) {
	defer wg.Done()
	result := sumList(nums)
	ch <- result
}

func splitList(list []int64, pieces int) [][]int64 {
	subListLength := (len(list)) / pieces
	resultList := make([][]int64, pieces)
	for i := 0; i < pieces; i++ {
		start := i * subListLength
		end := start + subListLength
		resultList[i] = list[start:end]
	}
	return resultList
}

func sumList(numbers []int64) int64 {
	var total int64
	for _, num := range numbers {
		total += num
	}
	return total
}

func generateNumList(length int) []int64 {
	nums := make([]int64, length)
	for i := 0; i < length; i++ {
		nums[i] = randomNumUnder(100)
	}
	return nums
}

func randomNumUnder(num int) int64 {
	return int64(rand.Int() % num)
}
