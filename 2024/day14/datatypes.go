package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	// "strings"
)

type Quadrant struct {
	TopLeft     int
	TopRight    int
	BottomLeft  int
	BottomRight int
}

func (quadrant Quadrant) safetyFactor() int {
	return quadrant.TopLeft * quadrant.TopRight * quadrant.BottomLeft * quadrant.BottomRight
}

type Direction struct {
	XOffset int
	YOffset int
}

type Robot struct {
	X         int
	Y         int
	MaxWidth  int
	MaxHeight int
	HasMoved  bool
	Velocity  Direction
}

func (robot *Robot) Move(timeInSeconds int) {
	xOffset := (robot.X + (robot.Velocity.XOffset * timeInSeconds))
	yOffset := (robot.Y + (robot.Velocity.YOffset * timeInSeconds))

	if xOffset < 0 {
		xOffset = robot.MaxWidth + (xOffset % robot.MaxWidth)
	}
	if yOffset < 0 {
		yOffset = robot.MaxHeight + (yOffset % robot.MaxHeight)
	}

	robot.X = xOffset % robot.MaxWidth
	robot.Y = yOffset % robot.MaxHeight
	robot.HasMoved = true
}

type RobotGroup []*Robot

type Bathroom [][]RobotGroup

func (room Bathroom) Display() {
	fmt.Println("Display Bathroom")
	for _, rowValue := range room {
		fmt.Println(rowValue)
	}
}

func (room Bathroom) FileDisplay() {
	file, err := os.OpenFile("testing.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error opening file:", err)
	}
	defer file.Close()
	var stringBuilder strings.Builder
	stringBuilder.WriteString("Display Bathroom\n")
	for _, rowValue := range room {
		for _, col := range rowValue {
			if len(col) > 0 {
				stringBuilder.WriteString("*")
			} else {
				stringBuilder.WriteString(" ")
			}
		}
		stringBuilder.WriteString("\n")
	}
	_, err = file.WriteString(stringBuilder.String())
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
func (room Bathroom) ResetMovementStatus() {
	for _, row := range room {
		for _, col := range row {
			for _, robot := range col {
				robot.HasMoved = false
			}
		}
	}
}
func (room Bathroom) MoveRobots(seconds int) {
	room.ResetMovementStatus()
	for Y, row := range room {
		for X, robotGroup := range row {
			for len(robotGroup) > 0 {
				robot := robotGroup[0]
				if robot.HasMoved {
					break
				}
				robot.Move(seconds)
				destinationRobotGroup := room[robot.Y][robot.X]
				room[robot.Y][robot.X] = append(destinationRobotGroup, robot)
				robotGroup = robotGroup[1:]
				room[Y][X] = robotGroup

			}
		}
	}
}

func (room Bathroom) getQuadrantCount() *Quadrant {
	maxWidth, maxHeight := len(room[0]), len(room)
	solutionQuadrant := &Quadrant{}
	midPointOfX := maxWidth / 2
	midPointOfY := maxHeight / 2
	for Y, row := range room {
		for X, col := range row {
			if Y < midPointOfY {
				if X < midPointOfX {
					solutionQuadrant.TopLeft += len(col)
					continue
				} else if X > midPointOfX {
					solutionQuadrant.TopRight += len(col)
					continue
				}

			} else if Y > midPointOfY {
				if X < midPointOfX {
					solutionQuadrant.BottomLeft += len(col)
					continue
				} else if X > midPointOfX {
					solutionQuadrant.BottomRight += len(col)
					continue
				}

			}

		}
	}
	return solutionQuadrant
}

func newBathroom(fileContent string, Width, Height int) Bathroom {
	bathRoom := Bathroom{}
	for range Height {
		newRow := []RobotGroup{}
		for range Width {
			robotGroup := RobotGroup{}
			newRow = append(newRow, robotGroup)

		}
		bathRoom = append(bathRoom, newRow)
	}
	pattern := `p=(-?\d+),(-?\d+) ?v=(-?\d+),(-?\d+)`
	reg := regexp.MustCompile(pattern)
	stringRobotData := reg.FindAllStringSubmatch(fileContent, -1)
	for _, robot := range stringRobotData {
		X := convertToInt(robot[1])
		Y := convertToInt(robot[2])
		robotGroup := bathRoom[Y][X]
		newRobot := &Robot{
			X:         X,
			Y:         Y,
			MaxWidth:  Width,
			MaxHeight: Height,
			Velocity:  Direction{convertToInt(robot[3]), convertToInt(robot[4])},
		}
		robotGroup = append(robotGroup, newRobot)
		bathRoom[Y][X] = robotGroup
	}

	return bathRoom
}

func convertToInt(str string) int {
	integer, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return integer
}
