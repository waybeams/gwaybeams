package control_test

import (
	"assert"
	"testing"
	"ui2/control"
)

func TestControl(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		instance := control.New()
		assert.NotNil(t, instance)
	})
}
