package display

import (
	"log"
	"os"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	// This is required if any test uses a OpenTestWindow
	runtime.LockOSThread()
	os.Exit(m.Run())
}

var TestComponent = NewComponentFactory("TestComponent", NewComponent)

// OpenTestWindow is a component option that will create and launch a new source with your
// component instance displayed inside of it. Any ComponentOptions provided to the call
// will be applied to the newly created source object.
func OpenTestWindow(userOptions ...ComponentOption) ComponentOption {
	return func(d Displayable) error {
		options := []ComponentOption{
			Width(800),
			Height(600),
			HAlign(AlignCenter),
			VAlign(AlignCenter),
			Children(func(b Builder, w Displayable) {
				w.AddChild(d)
			}),
		}
		options = append(options, userOptions...)

		win, err := NanoWindow(NewBuilder(), options...)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		win.(Window).Init()
		return nil
	}
}
