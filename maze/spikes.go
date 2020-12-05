package maze

type spikes struct {
	spikes []coordinates
}

func newSpikes(s []coordinates) *spikes {
	return &spikes{
		spikes:	s
	}
}
