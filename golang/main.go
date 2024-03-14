package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	amountOfRandomNumbers, numberOfGroups := getInput()
	numberList := generateNumList(amountOfRandomNumbers)
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

func getInput() (int, int) {
	var input string
	fmt.Println("Please enter the amount of random numbers you want to generate: ")
	fmt.Scan(&input)
	amountOfRandomNumbers, err := strconv.Atoi(input)
	for {
		if err == nil {
			break
		}
		fmt.Println("Invalid input. Please enter an integer: ")
		fmt.Scan(&input)
		amountOfRandomNumbers, err = strconv.Atoi(input)
	}
	numberOfGroups := len(input)
	return amountOfRandomNumbers, numberOfGroups
}

func beginSumming(nums []int64, ch chan int64, wg *sync.WaitGroup) {
	defer wg.Done()
	result := sumList(nums)
	ch <- result
}

func splitList(list []int64, pieces int) [][]int64 {
	fmt.Printf("Splitting list into %d pieces\n", pieces)
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

func concurrentExample() {
    // Start a goroutine to perform a concurrent task
    go func() {
        for i := 0; i < 5; i++ {
            fmt.Println("Goroutine: ", i)
            time.Sleep(time.Second)
        }
    }()
    // Main goroutine continues execution
    for i := 0; i < 5; i++ {
        fmt.Println("Main: ", i)
        time.Sleep(time.Second)
    }
}

func secondConcurrentExample() {
	// Main goroutine
	fmt.Println("Main goroutine starts")

	// Start a new goroutine
	go func() {
		fmt.Println("Goroutine starts")
		time.Sleep(time.Second)
		fmt.Println("Goroutine ends")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Main goroutine ends")
}