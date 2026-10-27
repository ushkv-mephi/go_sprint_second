package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements returns a slice of random integers.
func generateRandomElements(size int) ([]int, error) {
	if size <= 0 {
		return nil, fmt.Errorf("invalid slice size: %d", size)
	}

	newSlice := make([]int, size)
	for i := range size {
		newSlice[i] = rand.Intn(size) + 1
	}

	return newSlice, nil
}

// maximum returns the largest element in a slice.
func maximum(data []int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("the length must be positive")
	}

	return slices.Max(data), nil
}

// maxChunks returns the maximum value found using parallel processing.
func maxChunks(data []int) (int, error) {
	dataLength := len(data)
	if dataLength == 0 {
		return 0, errors.New("the length must be positive")
	}

	if dataLength <= CHUNKS {
		return slices.Max(data), nil
	}

	var wg sync.WaitGroup
	maxValues := make([]int, CHUNKS)
	chunkLength := dataLength / CHUNKS

	for i := range CHUNKS {
		startIndex := i * chunkLength
		endIndex := startIndex + chunkLength
		if i == CHUNKS-1 {
			endIndex = dataLength
		}
		chunk := data[startIndex:endIndex]
		wg.Add(1)

		go func(c []int) {
			defer wg.Done()
			chunkMax := slices.Max(c)
			maxValues[i] = chunkMax
		}(chunk)
	}

	wg.Wait()
	return slices.Max(maxValues), nil
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	elements, err := generateRandomElements(SIZE)
	if err != nil {
		log.Fatalf("Critical error: %v", err)
	}

	fmt.Println("Ищем максимальное значение в один поток")

	startTime := time.Now()
	max, err := maximum(elements)
	if err != nil {
		log.Fatalf("Critical error: %v", err)
	}

	elapsed := time.Since(startTime).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	startTime = time.Now()
	max, err = maxChunks(elements)
	if err != nil {
		log.Fatalf("Critical error: %v", err)
	}

	elapsed = time.Since(startTime).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", max, elapsed)
}
