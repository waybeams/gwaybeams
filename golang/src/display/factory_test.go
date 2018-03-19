package display

import (
	"assert"
	"regexp"
	"testing"
)

func TestFactory(t *testing.T) {
	instance := &Factory{}

	t.Run("Instantiable", func(t *testing.T) {
		assert.NotNil(instance)
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

	t.Run("Fails on too many arguments", func(t *testing.T) {
		one := &Opts{}
		two := &Opts{}
		three := &Opts{}
		four := &Opts{}
		args := []interface{}{one, two, three, four}
		_, err := ProcessArgs(args)

		if err == nil {
			t.Error("Expected ProcessArgs failure with too many arguments")
		} else {
			errStr := err.Error()
			matched, _ := regexp.MatchString("Too many arguments", errStr)
			if !matched {
				t.Errorf("Unexpected error message: %v", errStr)
			}
		}
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

	t.Run("Fails with Compose AND ComposeWithUpdate", func(t *testing.T) {
		one := func() {}
		two := func(update func()) {}
		args := []interface{}{one, two}
		_, err := ProcessArgs(args)
		if err == nil {
			t.Error("Expected error return")
		}
	})
}
