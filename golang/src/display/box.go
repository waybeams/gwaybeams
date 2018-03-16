package display

type box struct {
	sprite
}

func NewBox() Composable {
	lastId++
	return &box{
		sprite{
			id: lastId,
		},
	}
}
