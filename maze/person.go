package maze

import (
	"errors"
)

type direction int

// Four directions
const (
	RIGHT direction = 1 + iota
	LEFT
	UP
	DOWN
)

type person struct {
	position coordinates
}

func newPerson(pos coordinates) *person {
	return &person{
		position: pos,
	}
}

func (p *person) die() error {
	return errors.New("Death")
}

func (p *person) move(d direction) error {

	
	switch d {
	case RIGHT:
		p.position.x++
	case LEFT:
		p.position.x--
	case UP:
		p.position.y++
	case DOWN:
		p.position.y--
	}
	return nil
}
