package controls

import (
	"os"
	"runtime"
	"testing"
	"ui"
	"ui/control"
	"uiold/opts"
)

var TestControl = control.Define("TestControl", control.New)

type FakeControl struct {
	control.Control
}

func NewFake() *FakeControl {
	return &FakeControl{}
}

// Create a new factory using our control creation function reference.
var Fake = control.Define("Fake",
	func() ui.Displayable { return NewFake() })

func TestMain(m *testing.M) {
	// This is required if any test uses a OpenTestWindow
	runtime.LockOSThread()
	os.Exit(m.Run())
}

// OpenTestWindow is a control option that will create and launch a new window with your
// control instance displayed inside of it. Any options provided to the call
// will be applied to the newly created window object.
func OpenTestWindow(userOptions ...ui.Option) ui.Option {
	return func(d ui.Displayable) {
		options := []ui.Option{
			opts.Width(800),
			opts.Height(600),
			opts.HAlign(ui.AlignCenter),
			opts.VAlign(ui.AlignCenter),
			opts.Children(func(c ui.Context, w ui.Displayable) {
				w.AddChild(d)
			}),
		}
		options = append(options, userOptions...)

		win := NanoWindow(d.Context(), options...)
		win.(*NanoWindowControl).Listen()
	}
}
