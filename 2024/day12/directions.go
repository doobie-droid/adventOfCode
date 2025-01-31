package main

type Direction struct {
	rowPosition int
	colPosition int
	Name        string
}

const UP = "UP"
const DOWN = "DOWN"
const LEFT = "LEFT"
const RIGHT = "RIGHT"

var upDirection Direction = Direction{-1, 0, UP}
var rightDirection Direction = Direction{0, 1, RIGHT}
var downDirection Direction = Direction{1, 0, DOWN}
var leftDirection Direction = Direction{0, -1, LEFT}
var directions []Direction = []Direction{upDirection, rightDirection, downDirection, leftDirection}
