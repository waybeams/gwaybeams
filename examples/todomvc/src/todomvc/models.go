package todomvc

import (
	"clock"
	"fmt"
	"strings"
	"time"
)

const DefaultFilterSelection = "All"

var lastId = 0

type TodoItemModel struct {
	CompletedAt time.Time
	CreatedAt   time.Time
	DeletedAt   time.Time
	ID          int
	Text        string
	container   *TodoAppModel
}

func (t *TodoItemModel) Complete() {
	t.container.CompleteItem(t)
}

type TodoAppModel struct {
	Clock           clock.Clock
	items           []*TodoItemModel
	filterSelection string
}

func (t *TodoAppModel) PendingItems() []*TodoItemModel {
	result := []*TodoItemModel{}
	for _, item := range t.items {
		if item.CompletedAt.IsZero() {
			result = append(result, item)
		}
	}
	return result
}

func (t *TodoAppModel) PendingLabel() string {
	count := len(t.PendingItems())
	switch count {
	case 0:
		return "0 items left"
	case 1:
		return "1 item left"
	default:
		return fmt.Sprintf("%v items left", count)
	}
}

func (t *TodoAppModel) CompletedItems() []*TodoItemModel {
	result := []*TodoItemModel{}
	for _, item := range t.items {
		if !item.CompletedAt.IsZero() {
			result = append(result, item)
		}
	}
	return result
}

func (t *TodoAppModel) AllItems() []*TodoItemModel {
	return t.items
}

func (t *TodoAppModel) LastItem() *TodoItemModel {
	return t.items[len(t.items)-1]
}

func (t *TodoAppModel) UpdateItemAt(index int, text string) {
	t.items[index].Text = strings.Trim(text, " ")
}

func (t *TodoAppModel) ItemAt(index int) *TodoItemModel {
	return t.items[index]
}

func (t *TodoAppModel) PushItem(text string) {
	t.items = append(t.items, &TodoItemModel{
		CreatedAt: time.Now(),
		Text:      strings.Trim(text, " "),
		container: t,
	})
}

func (t *TodoAppModel) CompleteItem(item *TodoItemModel) {
	item.CompletedAt = t.Clock.Now()
}

func (t *TodoAppModel) SetFilterSelection(value string) {
	t.filterSelection = value
}

func (t *TodoAppModel) FilterSelection() string {
	if t.filterSelection == "" {
		t.filterSelection = DefaultFilterSelection
	}
	return t.filterSelection
}
