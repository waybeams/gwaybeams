package spec_test

import (
	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/ctrl"
	"github.com/waybeams/waybeams/pkg/events"
	"github.com/waybeams/waybeams/pkg/opts"
	"testing"
)

func TestStateful(t *testing.T) {
	t.Run("Configurable", func(t *testing.T) {
		box := ctrl.Box(
			opts.OnState("abcd", opts.Text("ABCD")),
			opts.OnState("efgh", opts.Text("EFGH")),
			opts.On("enter-abcd", opts.OptionsHandler(opts.SetState("abcd"))),
			opts.On("enter-efgh", opts.OptionsHandler(opts.SetState("efgh"))),
		)

		// Defaults to first declared state.
		assert.Equal(box.State(), "abcd")
		box.Emit(events.New("enter-efgh", box, nil))
		assert.Equal(box.State(), "efgh")
	})
}
