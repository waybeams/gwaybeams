package todomvc

import (
	"assert"
	"clock"
	"ctx"
	"opts"
	"testing"
)

func TestCreate(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		app := Create(clock.NewFake(), &TodoAppModel{})
		assert.NotNil(t, app)
	})

	t.Run("Todo Component", func(t *testing.T) {
		item := Todo(ctx.New(), opts.Data(&TodoItemModel{Text: "abcd"}))
		assert.Equal(t, item.Data().(*TodoItemModel).Text, "abcd")
	})
}
