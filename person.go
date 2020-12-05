package maze

import (
	"errors"
)

type direction int

const (
	RIGHT direction = 1 + iota
	LEFT
	UP
	DOWN
)

type person struct {
	position coordinates
}

func newPerson(pos []coordinates) *person {
	return &person {
		position: pos
	}
}

func (p *person) die() error {
	return errors.New("Death")
}

func (p *person) move(d direction) error {
	switch d {
	case RIGHT:
		p.pos.x++
	case LEFT:
		p.pos.x--
	case UP:
		p.pos.y++
	case DOWN:
		p.pos.y--
	}
}
