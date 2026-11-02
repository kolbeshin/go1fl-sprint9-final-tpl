package main

import (
	"fmt"
	mRand "math/rand"
	"sync"
	"time"
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

	result := make([]int, size)

	for i := range size {
		result[i] = mRand.Intn(100)
	}

	return result
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	vmax := data[0]

	for _, v := range data[1:] {
		if v > vmax {
			vmax = v
		}
	}
	return vmax
}

// splitIntoChunks необходима для разделения массива на части,
// чтобы предусмотреть корректное распределение данных между горутинами,
// особенно если количество элементов не делится ровно на количсетво CHUNKS
func splitIntoChunks(length int) []Chunk {
	if length <= 0 {
		return nil
	}

	count := min(CHUNKS, length)
	chunkSize := length / count
	overflow := length % count
	chunks := make([]Chunk, 0, count)
	start := 0

	for i := range count {
		size := chunkSize
		if i < overflow {
			size++
		}
		end := start + size
		chunks = append(chunks, Chunk{
			Start: start,
			End:   end,
		})
		start = end
	}
	return chunks
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunks := splitIntoChunks(len(data))
	maxValues := make([]int, len(chunks))

	var wg sync.WaitGroup

	for i, chunk := range chunks {
		wg.Add(1)

		go func() {
			defer wg.Done()

			part := data[chunk.Start:chunk.End]
			vmax := maximum(part)

			maxValues[i] = vmax
		}()
	}

	wg.Wait()

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	numbers := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxValue := maximum(numbers)
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d us\n", maxValue, elapsed.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxValue = maxChunks(numbers)
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d us\n", maxValue, elapsed.Microseconds())
}
