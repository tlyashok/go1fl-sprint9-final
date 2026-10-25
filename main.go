package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8

	// maxRandomValue — верхняя граница (включительно) для генерируемых чисел.
	maxRandomValue = 1_000_000_000
)

// generateRandomElements возвращает слайс заданного размера, заполненный
// случайными положительными числами из диапазона [1, maxRandomValue].
// Для неположительного размера возвращает пустой слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(maxRandomValue) + 1
	}

	return data
}

// maximum возвращает максимальное число в слайсе, поиск идёт в один поток.
// Для пустого слайса возвращает 0: все генерируемые числа положительные,
// поэтому 0 не может оказаться настоящим максимумом.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks возвращает максимальное число в слайсе, разбивая его на CHUNKS частей.
// Максимум каждой части ищется в отдельной горутине, затем среди найденных
// максимумов выбирается наибольший.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	// Слайс короче CHUNKS нельзя разбить на непустые части.
	if chunkSize == 0 {
		return maximum(data)
	}

	// Каждая горутина пишет только в свой индекс, поэтому мьютекс не нужен.
	chunkMaximums := make([]int, CHUNKS)

	var wg sync.WaitGroup
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		// Последняя часть забирает остаток, если длина слайса не делится на CHUNKS.
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(index int, chunk []int) {
			defer wg.Done()
			chunkMaximums[index] = maximum(chunk)
		}(i, data[start:end])
	}
	wg.Wait()

	return maximum(chunkMaximums)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxValue := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxValue = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)
}
