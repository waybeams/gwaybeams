package display

import (
	"assert"
	"testing"
)

/*
func FakeComponent(opts *Opts) {
	decl, err := CreateDeclaration(args)

	if err != nil {
		panic(err)
	}

	// Instantiate and configure the component
	sprite := NewSprite()
	sprite.Declaration(decl)
	f.Push(sprite)
}

func FakeRender(f Factory) {
	FakeComponent(f, &Opts{Id: "root"}, func() {
		FakeComponent(f, &Opts{Id: "child1", BackgroundColor: 0xfc0})
		FakeComponent(f)
	})
}
*/

func TestFactory(t *testing.T) {
	instance := NewFactory()

	t.Run("Instantiable", func(t *testing.T) {
		assert.NotNil(instance)
	})

	/*
		t.Run("GetRoot returns first component", func(t *testing.T) {
			t.Run("is callable", func(t *testing.T) {
				f := NewFactory()
				FakeRender(f)
				root := f.GetRoot()
				assert.NotNil(root)
				// assert.Equal(root.Id(), "efgh")
				assert.Equal(reflect.TypeOf(root).String(), "*display.Sprite")
			})
		})
	*/

	t.Run("Forwards stack.Push(nil) error", func(t *testing.T) {
		err := instance.Push(nil)
		assert.NotNil(err)
	})
}
