package main

func getPriceOfFencePartTwo(fileName string) int {
	// get string content from file
	fileContent := readFile(fileName)

	//read the string content into a matrix of rows and columns
	farm := getMatrix(fileContent)

	farm.locateDistinctSides()

	// assign the distinct sides to sides so the area can be calculated easily
	for _, rowValue := range farm {
		for _, plot := range rowValue {
			plot.Sides = plot.DistinctSides
		}
	}

	area := farm.calculateArea()
	return area
}
