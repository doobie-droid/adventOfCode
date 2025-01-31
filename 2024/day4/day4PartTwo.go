package main

//Algorithm is look out for all the A's
// if you find one check that the positions of all its diagonals are valid, if they are not, there is no mas
// check that m faces s at a diagonal and m does not face m or s does not face s, if the diagonals are equal,
// there is no mas
// while checking each diagonal, delete the seen character from available characters to ensure
// that you only see 2 S' and 2 M's , if you see any more, there is no MAS

func (point *Coordinate) twoMasExist(board [][]*Coordinate) bool {
	hashMapOfDiagonalLetters := map[rune]int{'S': 2, 'M': 2}
	movementList := []struct {
		rowMovement int
		colMovement int
	}{
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}
	for _, move := range movementList {
		rowPosition := point.row + move.rowMovement
		inverseRowPosition := point.row - move.rowMovement
		colPosition := point.col + move.colMovement
		inverseColPosition := point.col - move.colMovement
		if !point.isValidIndex(rowPosition, colPosition) || !point.isValidIndex(inverseRowPosition, inverseColPosition) {
			return false
		}
		if board[rowPosition][colPosition].Value == board[inverseRowPosition][inverseColPosition].Value {
			return false
		}
		charAtDiagonal := board[rowPosition][colPosition].Value
		if hashMapOfDiagonalLetters[charAtDiagonal] <= 0 {
			return false
		}
		hashMapOfDiagonalLetters[charAtDiagonal]--
	}
	return true
}

func findTwoMASInXShapeInFile(fileName string) int {
	xmasCount := 0
	matrix := readStringsIntoRuneMatrix(fileName)
	for _, rowValue := range matrix {
		for _, point := range rowValue {
			if point.Value != 'A' {
				continue
			}
			if point.twoMasExist(matrix) {
				xmasCount++
			}
		}
	}
	return xmasCount
}
