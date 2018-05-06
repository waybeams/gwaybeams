package todomvc

import (
	"assert"
	"clock"
	"ctx"
	"opts"
	"testing"
	"ui"
)

func createTodoTestApp() ui.Displayable {
	return Create(&TodoAppModel{},
		ctx.Clock(clock.NewFake()),
		ctx.Font("Roboto", "../../../../third_party/fonts/Roboto/Roboto-Regular.ttf"),
		ctx.Font("Roboto-Thin", "../../../../third_party/fonts/Roboto/Roboto-Thin.ttf"),
		ctx.Font("Roboto-Light", "../../../../third_party/fonts/Roboto/Roboto-Light.ttf"),
		ctx.Font("Roboto-Bold", "../../../../third_party/fonts/Roboto/Roboto-Bold.ttf"),
	)

}

func TestCreate(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		app := createTodoTestApp()
		assert.NotNil(t, app)
	})

	t.Run("Todo Component", func(t *testing.T) {
		item := Todo(ctx.New(), opts.Data(&TodoItemModel{Text: "abcd"}))
		assert.Equal(t, item.Data().(*TodoItemModel).Text, "abcd")
	})
}
