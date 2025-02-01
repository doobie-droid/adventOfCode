package main

type Plot struct {
	Row           int
	Col           int
	RowBounds     int
	ColBounds     int
	Value         string
	Matrix        Matrix
	Region        string
	Sides         map[string]bool
	DistinctSides map[string]bool
}

func (plot *Plot) hasBoundary(direction Direction) bool {
	if plot.Row+direction.rowOffset < 0 || plot.Col+direction.colOffset < 0 {
		return true
	}
	if plot.Row+direction.rowOffset >= plot.RowBounds || plot.Col+direction.colOffset >= plot.ColBounds {
		return true
	}
	return false
}

func (plot *Plot) CalculatePerimeter() int {
	numberOfValidSides := 0
	for _, sideExists := range plot.Sides {
		if sideExists {
			numberOfValidSides++
		}
	}
	return numberOfValidSides
}
func (plot *Plot) AddPlotToRegion(regionName string) {
	// if the plant is already in the same region, no need to bother adding it
	if plot.Region == regionName {
		return
	}

	// add the plot to the region passed along
	plot.Region = regionName

	// loop through all available directions
	for _, direction := range directions {
		// if a certain direction has a boundary, it means there is no plot
		// in that direction , skip the iteration
		if plot.hasBoundary(direction) {
			continue
		}
		// get the next plot
		nextRow := plot.Row + direction.rowOffset
		nextCol := plot.Col + direction.colOffset
		nextPlot := plot.Matrix[nextRow][nextCol]
		// if the next plot is equal to the current plot, move to the next plot and add repeat
		if plot.Value == nextPlot.Value {
			nextPlot.AddPlotToRegion(regionName)
		}
	}
}
