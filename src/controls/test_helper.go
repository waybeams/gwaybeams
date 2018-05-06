package controls

import (
	"component"
	"ctx"
	"opts"
	"os"
	"runtime"
	"testing"
	"ui"
)

var TestComponent = component.Define("TestComponent", component.New)

type FakeComponent struct {
	component.Component
}

func NewFake() *FakeComponent {
	return &FakeComponent{}
}

// Create a new factory using our component creation function reference.
var Fake = component.Define("Fake",
	func() ui.Displayable { return NewFake() })

func TestMain(m *testing.M) {
	// This is required if any test uses a OpenTestWindow
	runtime.LockOSThread()
	os.Exit(m.Run())
}

// OpenTestWindow is a component option that will create and launch a new window with your
// component instance displayed inside of it. Any ComponentOptions provided to the call
// will be applied to the newly created window object.
func OpenTestWindow(userOptions ...ui.Option) ui.Option {
	return func(d ui.Displayable) error {
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

		win := NanoWindow(ctx.New(), options...)
		win.(*NanoWindowComponent).Listen()
		return nil
	}
}
