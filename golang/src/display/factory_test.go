package display

import (
	"assert"
	"reflect"
	"testing"
)

func FakeComponent(f Factory, args ...interface{}) {
	decl, err := ProcessArgs(args)

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

func TestFactory(t *testing.T) {
	instance := NewFactory()

	t.Run("Instantiable", func(t *testing.T) {
		assert.NotNil(instance)
	})

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

	t.Run("Forwards stack.Push(nil) error", func(t *testing.T) {
		err := instance.Push(nil)
		assert.NotNil(err)
	})

	t.Run("Processes empty args", func(t *testing.T) {
		emptyArgs := []interface{}{}
		assert.NotNil(emptyArgs)
		decl, err := ProcessArgs(emptyArgs)
		assert.Nil(err)
		// assert.Nil does not work with interface{} args...
		// TODO(lbayes): Find a simple assertion library and use that instead.
		if decl.Compose != nil {
			t.Error("Expected Compose to be nil")
		}
		if decl.ComposeWithUpdate != nil {
			t.Error("Expected ComposeWithUpdate to be nil")
		}
	})

	t.Run("Processes provided Opts", func(t *testing.T) {
		color := uint(0xfc0)
		args := []interface{}{&Opts{BackgroundColor: color}}
		decl, _ := ProcessArgs(args)
		assert.Equal(decl.Options.BackgroundColor, color)
	})

	t.Run("Processes Compose (and not ComposeWithUpdate)", func(t *testing.T) {
		childrenFunc := func() {}
		args := []interface{}{childrenFunc}
		decl, _ := ProcessArgs(args)
		if decl.Compose == nil {
			t.Error("Expected Compose assignment")
		}

		if decl.ComposeWithUpdate != nil {
			t.Error("Expected ComposeWithUpdate assignment")
		}
	})

	t.Run("Process Compose with update (and not Compose)", func(t *testing.T) {
		childrenFunc := func(update func()) {}
		args := []interface{}{childrenFunc}
		decl, _ := ProcessArgs(args)
		if decl.ComposeWithUpdate == nil {
			t.Error("Expected ComposeWithUpdate assignment")
		}
		if decl.Compose != nil {
			t.Error("Expected Compose to be nil")
		}
	})

	t.Run("Errors", func(t *testing.T) {
		t.Run("Fails on too many arguments", func(t *testing.T) {
			args := []interface{}{&Opts{}, &Opts{}, &Opts{}, &Opts{}}
			_, err := ProcessArgs(args)

			assert.ErrorMatches("Too many arguments", err)
		})

		t.Run("Fails with multiple Opts", func(t *testing.T) {
			args := []interface{}{&Opts{}, &Opts{}}
			_, err := ProcessArgs(args)

			assert.ErrorMatches("Only one Opts", err)
		})

		t.Run("Fails with multiple func()", func(t *testing.T) {
			args := []interface{}{func() {}, func() {}}
			_, err := ProcessArgs(args)

			assert.ErrorMatches("Only one Compose function", err)
		})

		t.Run("Fails with multiple func(func())", func(t *testing.T) {
			one := func(update func()) {}
			two := func(update func()) {}
			args := []interface{}{one, two}
			_, err := ProcessArgs(args)

			assert.ErrorMatches("Only one ComposeWithUpdate", err)
		})

		t.Run("Fails with Compose AND ComposeWithUpdate", func(t *testing.T) {
			one := func() {}
			two := func(update func()) {}
			args := []interface{}{one, two}
			_, err := ProcessArgs(args)

			assert.ErrorMatches("Only one composition function", err)
		})
	})
}
