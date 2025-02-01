package main

type Direction struct {
	rowOffset int
	colOffset int
	Name      string
}

const UP, DOWN, LEFT, RIGHT = "UP", "DOWN", "LEFT", "RIGHT"

var upDirection, rightDirection, downDirection, leftDirection Direction = Direction{-1, 0, UP}, Direction{0, 1, RIGHT}, Direction{1, 0, DOWN}, Direction{0, -1, LEFT}

// List of 4 possible directions so that you can easily iterate
var directions []Direction = []Direction{upDirection, rightDirection, downDirection, leftDirection}
