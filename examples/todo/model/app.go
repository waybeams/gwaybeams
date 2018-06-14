package model

import (
	"time"
)

type ItemsShown int

const (
	AllItems = iota
	ActiveItems
	CompletedItems
)

type App struct {
	showing     ItemsShown
	allItems    []*Item
	enteredText string
}

func (t *App) Showing() ItemsShown {
	return t.showing
}

func (t *App) ClearCompleted() {
	t.allItems = t.ActiveItems()
}

func (t *App) ShowActiveItems() {
	t.showing = ActiveItems
}

func (t *App) ShowCompletedItems() {
	t.showing = CompletedItems
}

func (t *App) ShowAllItems() {
	t.showing = AllItems
}

func (t *App) DeleteItem(deletedItem *Item) {
	result := []*Item{}
	// Splice instead of rebuild!
	for _, item := range t.AllItems() {
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

func (t *App) CreateItem(desc string) *Item {
	item := &Item{
		Description: desc,
		CreatedAt:   time.Now(),
		collection:  t,
	}

	t.allItems = append(t.AllItems(), item)
	return item
}

func (t *App) AllItems() []*Item {
	return t.allItems
}

func (t *App) CurrentItems() []*Item {
	switch t.showing {
	case AllItems:
		return t.AllItems()
	case ActiveItems:
		return t.ActiveItems()
	case CompletedItems:
		return t.CompletedItems()
	}
	panic("CurrentItems was not configured properly")
	return nil
}

func (t *App) CompletedItems() []*Item {
	result := []*Item{}
	for _, item := range t.AllItems() {
		if !item.CompletedAt.IsZero() {
			result = append(result, item)
		}
	}
	return result
}

func (t *App) ActiveItems() []*Item {
	result := []*Item{}
	for _, item := range t.AllItems() {
		if item.CompletedAt.IsZero() {
			result = append(result, item)
		}
	}
	return result
}

func New() *App {
	return &App{}
}

func NewSample() *App {
	m := New()
	m.CreateItem("Item One")
	m.CreateItem("Item Two")
	m.CreateItem("Item Three")
	m.CreateItem("Item Four")
	m.CreateItem("Item Five")
	m.CreateItem("Item Six")
	return m
}
