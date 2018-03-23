package display

import (
	"assert"
	"testing"
)

type fakeData struct{}

func TestOpts(t *testing.T) {
	t.Run("InitializeOpts", func(t *testing.T) {
		opts := &ComponentModel{Padding: 4}
		InitializeOpts(opts)
		assert.Equal(opts.PaddingLeft, -1.0)
	})

	t.Run("Padding values must be greater than zero", func(t *testing.T) {
		opts := &ComponentModel{PaddingLeft: -1}
		_, err := InitializeOpts(opts)
		assert.NotNil(err)
		assert.ErrorMatch("greater than zero", err)
	})

	t.Run("Processes empty args", func(t *testing.T) {
		emptyArgs := []interface{}{}
		assert.NotNil(emptyArgs)
		decl, err := NewDeclaration(emptyArgs)
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

	t.Run("Processes provided ComponentModel", func(t *testing.T) {
		args := []interface{}{&ComponentModel{Disabled: true}}
		decl, _ := NewDeclaration(args)
		assert.Equal(decl.Options.Disabled, true)
	})

	t.Run("Processes Compose (and not ComposeWithUpdate)", func(t *testing.T) {
		childrenFunc := func() {}
		args := []interface{}{childrenFunc}
		decl, _ := NewDeclaration(args)
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
		decl, _ := NewDeclaration(args)
		if decl.ComposeWithUpdate == nil {
			t.Error("Expected ComposeWithUpdate assignment")
		}
		if decl.Compose != nil {
			t.Error("Expected Compose to be nil")
		}
	})

	t.Run("Errors", func(t *testing.T) {
		t.Run("Fails on too many arguments", func(t *testing.T) {
			args := []interface{}{&ComponentModel{}, &ComponentModel{}, &ComponentModel{}, &ComponentModel{}}
			_, err := NewDeclaration(args)

			assert.ErrorMatch("Too many arguments", err)
		})

		t.Run("Fails with multiple ComponentModel", func(t *testing.T) {
			args := []interface{}{&ComponentModel{}, &ComponentModel{}}
			_, err := NewDeclaration(args)

			assert.ErrorMatch("Only one ComponentModel", err)
		})

		t.Run("Fails with multiple func()", func(t *testing.T) {
			args := []interface{}{func() {}, func() {}}
			_, err := NewDeclaration(args)

			assert.ErrorMatch("Only one Compose function", err)
		})

		t.Run("Fails with multiple func(func())", func(t *testing.T) {
			one := func(update func()) {}
			two := func(update func()) {}
			args := []interface{}{one, two}
			_, err := NewDeclaration(args)

			assert.ErrorMatch("Only one ComposeWithUpdate", err)
		})

		t.Run("Fails with Compose AND ComposeWithUpdate", func(t *testing.T) {
			one := func() {}
			two := func(update func()) {}
			args := []interface{}{one, two}
			_, err := NewDeclaration(args)

			assert.ErrorMatch("Only one composition function", err)
		})

		t.Run("Fails with multiple component Data", func(t *testing.T) {
			one := &fakeData{}
			two := &fakeData{}
			args := []interface{}{one, two}
			_, err := NewDeclaration(args)

			assert.ErrorMatch("Only one bag of component data", err)
		})
	})
}
