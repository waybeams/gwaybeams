# env/browser

This package provides support for running Waybeams applications in a modern browser.

Unfortunately, this package is not currently testable as the Gopherjs implementation consistently returns concrete types that only run error-free when connected to compiler-generated global state that only exists in the browser runtime environment.

I'm concerned about wrapping the entire Gopherjs implementation with more flexibly designed concrete adapters as I'm already working with adapters ([surface](https://github.com/waybeams/waybeams/blob/master/pkg/spec/surface.go), [clock](https://github.com/waybeams/waybeams/blob/master/pkg/clock/clock.go) and [window](https://github.com/waybeams/waybeams/blob/master/pkg/spec/window.go)), and an even larger secondary adapter could significantly bloat the browser payload size. I'm confident that I could continue to hack something together that mostly works, but since I cannot run automated tests against this environment, I'd have a difficult time making committments about it's reliability.

This work is made so much more difficult by the underlying implementation's reliance on concrete types.

Now that I've given up on an in-process adapter, I'm considering an approach that might involve forking or updating the Gopherjs compiler in such a way that the expected browser environment can be more easily mocked in a Go test environment.

Any tips or thoughts are welcome.

Exploration is ongoing.
