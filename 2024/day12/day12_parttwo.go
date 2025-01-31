package main

func getPriceOfFencePartTwo(fileName string) int {
	// get string content from file
	fileContent := readFile(fileName)

	//read the string content into a matrix of rows and columns
	_ = getMatrix(fileContent)
	return 4
}
