package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{name: "нулевой размер", size: 0, wantLen: 0},
		{name: "отрицательный размер", size: -10, wantLen: 0},
		{name: "один элемент", size: 1, wantLen: 1},
		{name: "много элементов", size: 10_000, wantLen: 10_000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)

			if data == nil {
				t.Fatalf("generateRandomElements(%d) вернула nil, ожидался слайс", tt.size)
			}
			if len(data) != tt.wantLen {
				t.Fatalf("generateRandomElements(%d): длина %d, ожидалась %d", tt.size, len(data), tt.wantLen)
			}
			for i, value := range data {
				if value < 1 || value > maxRandomValue {
					t.Fatalf("элемент %d = %d вне диапазона [1, %d]", i, value, maxRandomValue)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil слайс", data: nil, want: 0},
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{42}, want: 42},
		{name: "максимум в начале", data: []int{9, 3, 5, 1}, want: 9},
		{name: "максимум в середине", data: []int{3, 9, 5, 1}, want: 9},
		{name: "максимум в конце", data: []int{3, 5, 1, 9}, want: 9},
		{name: "все элементы одинаковые", data: []int{7, 7, 7, 7}, want: 7},
		{name: "несколько максимумов", data: []int{1, 9, 4, 9, 2}, want: 9},
		{name: "отрицательные числа", data: []int{-5, -1, -9}, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Fatalf("maximum(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil слайс", data: nil, want: 0},
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{42}, want: 42},
		{name: "меньше элементов, чем частей", data: []int{3, 8, 1}, want: 8},
		{name: "максимум в последнем элементе остатка", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100}, want: 100},
		{name: "максимум в первом элементе", data: []int{100, 2, 3, 4, 5, 6, 7, 8, 9, 10}, want: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Fatalf("maxChunks(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	sizes := []int{1, CHUNKS - 1, CHUNKS, CHUNKS + 1, 1_001, 100_000}

	for _, size := range sizes {
		data := generateRandomElements(size)
		want := maximum(data)

		if got := maxChunks(data); got != want {
			t.Fatalf("размер %d: maxChunks = %d, maximum = %d", size, got, want)
		}
	}
}
