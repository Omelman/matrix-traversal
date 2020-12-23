package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMainPart(t *testing.T) {
	var tests = []struct {
		inputMatrix   [][]int
		expectedArray []int
	}{
		{[][]int{
			{1},
		}, []int{1}},

		{[][]int{
			{1, 2},
			{3, 4},
		}, []int{1, 2, 4, 3}},

		{[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},

		{[][]int{
			{1, 2, 3, 1},
			{4, 5, 6, 4},
			{7, 8, 9, 7},
			{7, 8, 9, 7},
		}, []int{1, 2, 3, 1, 4, 7, 7, 9, 8, 7, 7, 4, 5, 6, 9, 8}},
	}

	for i, test := range tests {
		assert.Equal(t, test.expectedArray, makeClockwise(test.inputMatrix),
			fmt.Sprintf("test case: %d", i+1))
	}
}
