package main

func isMatrixSquare(matrix [][]int) bool {
	firstLen := len(matrix[0])
	for i := range matrix {
		if firstLen != len(matrix[i]) {
			return false
		}
	}
	return true
}

func makeClockwise(squareMatrix [][]int) []int {
	var (
		result    []int
		length    = len(squareMatrix)
		column    = 0
		row       = 0
		shift     = 0
		direction = 1
	)

	for len(result) < length*length {
		if shift == 0 {
			column = column - 1
		}
		//horizontal iteration
		for i := 0; i < length-shift; i++ {
			column = column + direction
			result = append(result, squareMatrix[row][column])
		}
		shift++
		//vertical iteration
		for i := 0; i < length-shift; i++ {
			row = row + direction
			result = append(result, squareMatrix[row][column])
		}
		direction = (-1) * direction
	}

	return result
}
