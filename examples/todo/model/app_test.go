package model_test

import (
	"github.com/waybeams/waybeams/examples/todo/model"
	"testing"
)

func TestAppModel(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		m := model.New()
		if m == nil {
			t.Error("Expected model instance")
		}
	})
}
