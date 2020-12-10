package maze

type arena struct {
	width  int
	height int
	person *person
	stairs *stairs
	spikes *spikes
}

func newArena(w, h int, p *person, s *stairs, sp *spikes) *arena {
	a := &arena{
		width:  w,
		height: h,
		person: p,
		stairs: s,
		spikes: sp,
	}
	return a
}

func (a *arena) personMove(d direction) error {
	if err := a.person.move(d); err != nil {
		return err
	}

	if a.personOnSpike() || a.personOnEdge() {
		a.person.die()
	}

	return nil
}

func (a *arena) personOnSpike() bool {
	for _, s := range a.spikes.position {
		if a.person.position.x == s.x && a.person.position.y == s.y {
			return true
		}
	}
	return false
}

func (a *arena) personOnEdge() bool {
	return a.person.position.x <= 0 || a.person.position.x >= a.width || a.person.position.y <= 0 || a.person.position.y >= a.height
}
