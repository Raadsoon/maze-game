package maze

type stairs struct {
	position coordinates
}

func newStairs(pos[] coordinates) *stairs {
	return &stairs{
		position: pos
	}
}