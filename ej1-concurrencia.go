package main

import (
	"fmt"
	"time"
)

func sumSequential() int {
	sum := 0
	for i := 1; i <= 100000000; i++ {
		sum += i
	}
	return sum
}

func sumConcurrent(start, end int, result chan<- int) {
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	result <- sum
}

func main() {
	// Suma secuencial
	startSequential := time.Now()
	resultSequential := sumSequential()
	elapsedSequential := time.Since(startSequential)

	fmt.Printf("Suma secuencial: %d\n", resultSequential)
	fmt.Printf("Tiempo transcurrido (secuencial): %s\n\n", elapsedSequential)

	// Suma concurrente
	startConcurrent := time.Now()

	resultChannel := make(chan int)

	middle := 50000000

	go sumConcurrent(1, middle, resultChannel)
	go sumConcurrent(middle+1, 100000000, resultChannel)

	partialSum1 := <-resultChannel
	partialSum2 := <-resultChannel

	totalSum := partialSum1 + partialSum2

	elapsedConcurrent := time.Since(startConcurrent)

	fmt.Printf("Suma concurrente: %d\n", totalSum)
	fmt.Printf("Tiempo transcurrido (concurrente): %s\n", elapsedConcurrent)
}
