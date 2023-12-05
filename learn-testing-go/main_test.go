package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	// Arrange
	testTable := []struct {
		numbers  []int
		expected int
	}{{
		numbers:  []int{76, 18, 5, 1, 57, 97, 81, 34, 49, 42, 10, 40, 90, 42, 39},
		expected: 97,
	}, {
		numbers:  []int{21, 43, 87, 84, 72, 8, 80, 52, 65, 52, 91, 55, 40, 48, 47},
		expected: 91,
	},
		{
			numbers:  []int{58, 62, 78, 29, 60, 49, 50, 39, 25, 53, 7, 93, 10, 87, 9},
			expected: 93,
		}, {
			numbers:  []int{69, 31, 8, 77, 83, 37, 28, 78, 88, 83, 25, 24, 36, 39, 91},
			expected: 91,
		},
		{
			numbers:  []int{70, 56, 13, 39, 79, 98, 32, 2, 73, 57, 94, 32, 12, 37, 81},
			expected: 98,
		},
		{
			numbers:  []int{46, 20, 86, 50, 45, 20, 9, 58, 64, 25, 21, 27, 55, 12, 38},
			expected: 86,
		},
	}

	// Act
	for _, testCase := range testTable {
		result := Max(testCase.numbers)
		t.Logf("Calling: Max(%v), result %d\n", testCase.numbers, result)
		// Assert
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrect result. Expected %d, got %d", testCase.expected, result))

	}
}
