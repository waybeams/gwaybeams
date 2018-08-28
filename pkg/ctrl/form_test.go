package ctrl_test

import (
	"testing"

	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/ctrl"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/opts"
	// "github.com/waybeams/waybeams/pkg/spec"
)

func TestForm(t *testing.T) {

	t.Run("Instantiable", func(t *testing.T) {
		var received events.Event = nil

		instance := ctrl.Form()
		instance.On(events.Submitted, func(e events.Event) {
			received = e
		})
		instance.Emit(events.New(events.EnterKeyReleased, instance, nil))

		m := received.Payload().(map[string]interface{})
		assert.Equal(received.Name(), events.Submitted)
		assert.Equal(len(m), 0)
	})

	t.Run("Provides child values", func(t *testing.T) {
		var received events.Event = nil

		instance := ctrl.Form(
			opts.Child(ctrl.TextInput(opts.Key("one"), opts.Text("abcd"))),
			opts.Child(ctrl.TextInput(opts.Key("two"), opts.Text("efgh"))),
		)
		instance.On(events.Submitted, func(e events.Event) {
			received = e
		})
		instance.Emit(events.New(events.EnterKeyReleased, instance, nil))

		m := received.Payload().(map[string]interface{})
		assert.Equal(received.Name(), events.Submitted)
		assert.Equal(len(m), 2)
		assert.Equal(m["one"], "abcd")
		assert.Equal(m["two"], "efgh")
	})
}
