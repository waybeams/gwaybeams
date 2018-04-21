package todomvc

import "time"

type TodoItemModel struct {
	CompletedAt time.Time
	CreatedAt   time.Time
	DeletedAt   time.Time
	Text        string
}

type TodoModel struct {
	Items []*TodoItemModel
}

func (t *TodoModel) PushItem(text string) {
	t.Items = append(t.Items, &TodoItemModel{
		Text:      text,
		CreatedAt: time.Now(),
	})
}
