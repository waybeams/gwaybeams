package todomvc

import (
	"github.com/waybeams/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		app := App(nil)
		assert.NotNil(t, app)
	})
}
