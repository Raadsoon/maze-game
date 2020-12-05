package maze

import (
	"github.com/nsf/termbox-go"
)

var (
	keyboardEventsChan = make(chan keyboardEvent)
)

type Game struct {
	arena		*arena
	isVisible	bool
	isOver		bool
}

func initialPerson() *person {
	return newPerson(
		coordinates{x: 1, y: 1}
	)
}

func initialStairs() *stairs {
	return newStairs(
		coordinates{x: 5, y: 5}
	)
}

func initialSpikes() *spikes {
	return newSpikes([]coordinates{
		coordinates{x: 2, y: 3}
		coordinates{x: 3, y: 3}
		coordinates{x: 4, y: 3}
		coordinates{x: 5, y: 3}
	})
}

func initialArena() *arena {
	return newArena(
		6, 6, initialPerson(), initialStairs(), initialSpikes()
	)
}

func (g *Game) quit() {
	g.isOver = true
}

func (g *Game) retry() {
	g.arena = initialArena()
	isOver = false
}

func newGame() *Game {
	return &Game{initialArena()}
}

func (g *Game) start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	defer termbox.Close()
	go listenToKeyboard(keyboardEventsChan)
	
	if err := g.render(); err != nil {
		panic(err)
	}

mainloop:
	for {
		select {
		case e := <- keyboardEventsChan:
			switch e.eventType {
			case MOVE:
				d = keyToDirection(e.key)
				g.arena.person.move(d)
				if err := g.render(); err != nil {
					panic(err)
			case RETRY:
				g.retry()
			case QUIT:
				break mainloop
		}
		default:
			if err := g.render(); err != nil {
				panic(err)
			}
	}
}