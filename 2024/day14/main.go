package main

import (
	// "fmt"
	"os"
)

func findSafetyFactor(fileName string, width, height int) int {
	lines := readFile(fileName)
	bathroom := newBathroom(lines, width, height)
	bathroom.MoveRobots(100)
	quadrant := bathroom.getQuadrantCount()
	return quadrant.safetyFactor()

}

// NOTE: THE PART TWO OF DAY 14 IS STILL UNDER CONSTRUCTION
// I HAVE NOT YET PUBLISHED A SOLUTION FOR THIS
func findSafetyFactorPartTwo(fileName string, width, height int) int {
	lines := readFile(fileName)
	bathroom := newBathroom(lines, width, height)
	for i := 0; i < 120; i++ {
		bathroom.MoveRobots(1)
		bathroom.FileDisplay()
	}
	quadrant := bathroom.getQuadrantCount()
	return quadrant.safetyFactor()
}

func main() {
	findSafetyFactorPartTwo("day14Sample.txt", 101, 103)
}

func readFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(file)
}
