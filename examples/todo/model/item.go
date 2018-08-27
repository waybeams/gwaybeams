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

func (t *Item) IsCompleted() bool {
	return !(t.CompletedAt == time.Time{})
}

func (t *Item) ToggleCompleted() {
	if t.CompletedAt.IsZero() {
		t.CompletedAt = time.Now()
	} else {
		t.CompletedAt = time.Time{}
	}
	if t.collection != nil {
		t.collection.OnItemChanged(t)
	}
}
