package todomvc

import (
	"assert"
	"clock"
	. "display"
	"testing"
)

func TestCreate(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		app, _ := Create(clock.NewFake(), &TodoAppModel{})
		assert.NotNil(t, app)
	})

	t.Run("Todo Component", func(t *testing.T) {
		item, _ := Todo(NewBuilder(), Data(&TodoItemModel{Text: "abcd"}))
		assert.Equal(t, item.Data().(*TodoItemModel).Text, "abcd")
	})
}
