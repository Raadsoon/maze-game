package maze

type spikes struct {
	position []coordinates
}

func newSpikes(s []coordinates) *spikes {
	return &spikes{
		position:	s,
	}
}
