package model

import "time"

type Item struct {
	CreatedAt   time.Time
	CompletedAt time.Time
	Description string
	collection  *App
}

func (t *Item) SetDescription(text string) {
	t.Description = text
}

func (t *Item) Delete() {
	t.collection.DeleteItem(t)
}

func (t *Item) ToggleCompleted() {
	if t.CompletedAt.IsZero() {
		t.CompletedAt = time.Now()
	} else {
		t.MarkActive()
	}
}

func (t *Item) MarkActive() {
	t.CompletedAt = time.Time{}
}
