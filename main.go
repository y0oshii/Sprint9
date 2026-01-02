package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	rand.NewSource(time.Now().Unix())

	result := make([]int, size)
	for i := 0; i < size; i ++ {
		result[i] = rand.Intn(SIZE)
	}

	return result
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	var ma int

	if len(data) == 0 {
		return 0
	}

	for _, i := range data {
		if i > ma{
			ma = i
		}
	}
	return ma
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	if chunkSize == 0 {
		return maximum(data)
	}

	maxValue := make([]int, CHUNKS)
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i ++ {
		start := i * chunkSize
		end := i + chunkSize

		go func(index, from, to int) {
			defer wg.Done()
			maxValue[index] = maximum(data[from:to]) 
		}(i, start, end)
	}
	wg.Wait()
	return maximum(maxValue)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
