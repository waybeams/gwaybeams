package display

type box struct {
	sprite
}

func NewBox() Displayable {
	lastId++
	return &box{
		sprite{
			id: lastId,
		},
	}
}
