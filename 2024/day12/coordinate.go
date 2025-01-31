package main

type Coordinate struct {
	Row       int
	Col       int
	RowBounds int
	ColBounds int
	Value     string
	Matrix    Matrix
	Perimeter int
	Region    string
}

func (point *Coordinate) hasBoundary(direction Direction) bool {
	if point.Row+direction.rowPosition < 0 || point.Col+direction.colPosition < 0 {
		return true
	}
	if point.Row+direction.rowPosition >= point.RowBounds || point.Col+direction.colPosition >= point.ColBounds {
		return true
	}
	return false
}

func (point *Coordinate) NameAllValidNeighbors(regionName string) {
	if point.Region == regionName {
		return
	}

	point.Region = regionName

	for _, direction := range directions {
		if point.hasBoundary(direction) {
			continue
		}
		nextRow := point.Row + direction.rowPosition
		nextCol := point.Col + direction.colPosition
		nextPoint := point.Matrix[nextRow][nextCol]
		if point.Value == nextPoint.Value {
			nextPoint.NameAllValidNeighbors(regionName)
		}
	}
}
