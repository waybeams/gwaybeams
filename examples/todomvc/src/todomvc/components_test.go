package todomvc

import (
	"assert"
	"clock"
	"testing"
	"ui"
	"ui/context"
	"ui/opts"
)

func createTodoTestApp() ui.Displayable {
	return Create(&TodoAppModel{},
		context.Clock(clock.NewFake()),
		context.Font("Roboto", "../../../../third_party/fonts/Roboto/Roboto-Regular.ttf"),
		context.Font("Roboto-Thin", "../../../../third_party/fonts/Roboto/Roboto-Thin.ttf"),
		context.Font("Roboto-Light", "../../../../third_party/fonts/Roboto/Roboto-Light.ttf"),
		context.Font("Roboto-Bold", "../../../../third_party/fonts/Roboto/Roboto-Bold.ttf"),
	)
}

func TestCreate(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		app := createTodoTestApp()
		assert.NotNil(t, app)
	})

	t.Run("Todo Control", func(t *testing.T) {
		item := Todo(context.New(), opts.Data("TodoModel", &TodoItemModel{Text: "abcd"}))
		assert.Equal(t, item.Data("TodoModel").(*TodoItemModel).Text, "abcd")
	})
}
