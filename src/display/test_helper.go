package display

import (
	"log"
	"os"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	// This is required if any test uses a TestWindow
	runtime.LockOSThread()
	os.Exit(m.Run())
}

var TestComponent = NewComponentFactory("TestComponent", NewComponent)

// TestWindow is a component option that will create and launch a new window with your
// component instance displayed inside of it. Any ComponentOptions provided to the call
// will be applied to the newly created window object.
func TestWindow(options ...ComponentOption) ComponentOption {
	return func(d Displayable) error {
		win, err := NanoWindow(NewBuilder(),
			Width(800),
			Height(600),
			HAlign(AlignCenter),
			VAlign(AlignCenter),
			Children(func(b Builder, w Displayable) {
				w.AddChild(d)
			}))
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		win.(Window).Init()
		return nil
	}
}
