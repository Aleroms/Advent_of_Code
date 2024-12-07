package main

type dir int
type point [2]int

const (
	north dir = iota
	east
	south
	west
)

type guard struct {
	direction        dir
	current_position point
}

type error interface {
	Error() string
}

// getDirection returns the new point that the guard will step towards
func (g guard) getDirection() point {
	var direction point
	switch g.direction {
	case north:
		direction = point{-1, 0}
	case east:
		direction = point{0, 1}
	case south:
		direction = point{1, 0}
	case west:
		direction = point{0, -1}
	}
	return direction
}

// headTowards updates the guards current position
func (g *guard) headTowards(direction point) {
	g.current_position[0] = direction[0]
	g.current_position[1] = direction[1]
}

// turnDirection turns 90 degrees to the right from current direction
func (g *guard) turnDirection() {
	switch g.direction {
	case north:
		g.direction = east
	case east:
		g.direction = south
	case south:
		g.direction = west
	case west:
		g.direction = north
	}
}

// comment test
type guardStartingInfo struct {
	direction dir
	position  point
}