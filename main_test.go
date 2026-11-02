package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "нулевой размер",
			size: 0,
		},
		{
			name: "отрицательный размер",
			size: -10,
		},
		{
			name: "один элемент",
			size: 1,
		},
		{
			name: "несколько элементов",
			size: 10,
		},
		{
			name: "большой слайс",
			size: 1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)

			expectedLen := max(0, tt.size)

			if len(result) != expectedLen {
				t.Fatalf(
					"Ожидалось %d, получено %d",
					expectedLen,
					len(result),
				)
			}

			for _, value := range result {
				if value < 0 || value >= 100 {
					t.Errorf(
						"Число %d выходит за диапазон [0, 100)",
						value,
					)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{
			name:     "пустой слайс",
			data:     []int{},
			expected: 0,
		},
		{
			name:     "nil слайс",
			data:     nil,
			expected: 0,
		},
		{
			name:     "один элемент",
			data:     []int{42},
			expected: 42,
		},
		{
			name:     "несколько элементов",
			data:     []int{10, 20, 5, 50, 30},
			expected: 50,
		},
		{
			name:     "отрицательные числа",
			data:     []int{-10, -5, -20},
			expected: -5,
		},
		{
			name:     "нули",
			data:     []int{0, 0, 0},
			expected: 0,
		},
		{
			name:     "одинаковые числа",
			data:     []int{7, 7, 7, 7},
			expected: 7,
		},
		{
			name:     "максимум в начале",
			data:     []int{100, 20, 30},
			expected: 100,
		},
		{
			name:     "максимум в конце",
			data:     []int{10, 20, 100},
			expected: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.data)

			if result != tt.expected {
				t.Errorf(
					"maximum(%v) = %d, ожидалось %d",
					tt.data,
					result,
					tt.expected,
				)
			}
		})
	}
}
