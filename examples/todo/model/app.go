package model

import "time"

const ShowAllItems = "ShowAllItems"
const ShowActiveItems = "ShowActiveItems"
const ShowCompletedItems = "ShowCompletedItems"

type App struct {
	CurrentListName string
	allItems        []*Item
	enteredText     string
}

func (t *App) ClearCompleted() {
	t.allItems = t.ActiveItems()
}

func (t *App) DeleteItem(deletedItem *Item) {
	result := []*Item{}
	// Splice instead of rebuild!
	for _, item := range t.allItems {
		if item != deletedItem {
			result = append(result, item)
		}
	}
	t.allItems = result
}

func (t *App) EnteredText() string {
	return t.enteredText
}

func (t *App) UpdateEnteredText(str string) {
	t.enteredText = str
}

func (t *App) CreateItemFromEnteredText() {
	t.CreateItem(t.EnteredText())
	t.enteredText = ""
}

func (t *App) CreateItem(desc string) {
	t.allItems = append(t.allItems, &Item{
		Description: desc,
		CreatedAt:   time.Now(),
		collection:  t,
	})
}

func (t *App) CurrentItems() []*Item {
	switch t.CurrentListName {
	case ShowAllItems:
		return t.allItems
	case ShowCompletedItems:
		return t.CompletedItems()
	default:
		return t.ActiveItems()
	}
}

func (t *App) CompletedItems() []*Item {
	result := []*Item{}
	for _, item := range t.allItems {
		if !item.CompletedAt.IsZero() {
			result = append(result, item)
		}
	}
	return result
}

func (t *App) ActiveItems() []*Item {
	result := []*Item{}
	for _, item := range t.allItems {
		if item.CompletedAt.IsZero() {
			result = append(result, item)
		}
	}
	return result
}

func New() *App {
	return &App{}
}
