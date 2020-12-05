package maze

import (
	"github.com/nsf/termbox-go"
)

const (
	defaultColor = termbox.ColorDefault
	bgColor		 = termbox.ColorDefault
	spikesColor	 = termbox.ColorGrey
	personColor	 = termbox.ColorYellow
	stairsColor	 = termbox.ColorOrange
)

func (g *Game) render() error {
	termbox.Clear(defaultColor, defaultColor)

	var (
		tbW, tbH = termbox.Size()
		left	 = (tbW - g.arena.width) / 2
		right	 = (tbW + g.arena.width) / 2
		bottom	 = (tbH + g.arena.height) / 2
		top		 = (tbH - g.arena.height) / 2
	)

	renderArena(top, bottom, left, right)
	renderPerson(g.arena.person, top, left)
	renderSpikes(g.arena.spikes, top, left)
	renderStairs(g.arena.stairs, top, left)

	return termbox.Flush()
}

func renderArena(a *arena, top, bottom, left, right) {
	for i := top; i <= bottom, i++ {
		termbox.SetCell(left, i, '|', defaultColor, bgColor)
		termbox.SetCell(left + a.width, i, '|', defaultColor, bgColor)
	}
	for i := left; i <= right; i++ {
		termbox.SetCell(i, top, '-', defaultColor, bgColor)
		termbox.SetCell(i, bottom, '-', defaultColor, bgColor)
	} 
	termbox.SetCell(left, top, '┌', defaultColor, bgColor)
	termbox.SetCell(left, bottom, '└', defaultColor, bgColor)
	termbox.SetCell(right, top, '┐', defaultColor, bgColor)
	termbox.SetCell(right, bottom, '┘', defaultColor, bgColor)
}

func renderPerson(p *person, top, left) {
	termbox.SetCell(left + p.x, top + p.y, 'P', personColor, personColor)
}

func renderSpikes(s *spikes, top, left) {
	for _, sp := range s {
		termbox.SetCell(left + sp.x, top + sp.y, 'X', spikesColor, spikesColor)
	}
}

func renderStairs(s *stairs, top, left)
	termbox.SetCell(left + s.x, top + s.y, 'O', stairsColor, stairsColor)
}