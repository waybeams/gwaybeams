
## This software is experimental. Contributions are welcome, but production use is discouraged

# Waybeams

With Waybeams you can quickly and reliably create and test tiny (<10MB). delightful applications that can thrive on a multitude of surfaces (Windows, macOS, Linux, Android, iOS, Raspberry Pi, Beaglebone, etc.).

![Waybeams Image](media/waybeams-home.jpg)

* Image provided courtesy of [Carlos Amato](https://www.flickr.com/photos/charlyamato/) and the [Creative Commons](https://creativecommons.org/licenses/by-nc-nd/2.0/) license(s).

_According to Merriam Webster, A Waybeam is, ": a beam supporting a way; specifically : either of two longitudinal beams resting on transverse girders and supporting the rails of a road crossing a bridge"_

This feels like an apt metaphor for a set of tools that make it possible to rapidly, safely and continuously build and deploy experiences.

Waybeams is built using the [Go language](https://golang.org/) and an OpenGL rendering surface (currently, [NanoVG](https://github.com/memononen/nanovg))

Waybeams provides:
* Simple, fast, reactive and composable GUI toolkit
* Cross-platform, GPU accelerated drawing surface
* Pure Go component declaration and configuration
* Tiny, blazing fast, constraints-based layouts
* Component Trait assignment using web-like selectors
* Headless (insanely fast) environment for UI tests
* Isolated visual environment for test-driven development on UI components (tbd)
* Automated image rendering surface (from tests) for release validation (tbd)

# Getting Started

Step one: Install some version of Go to the system > 1.4. This version will be used purely to build our local version.

Once this is done, change to this directory and run the following (on Darwin/macOS or Linux):

```bash
source setup-env.sh
make dev-install
```

# What is, "pure Go component declaration?"
Rather than using some separate language (or many, many more) to describe a user interface, we use _one_. We use pure Go to describe behavior, style _and_ structure. A core thesis of this work, is that this decision alone can deliver significant reductions in cognitive load, development time and even runtime performance.

A very simple Waybeams application might look something like the following:
```go
package main

import (
  . "display"
  "runtime"
)

func init() {
  runtime.LockOSThread()
}

func createWindow() (Displayable, error) {
  var messages := []string{"Hello World", "Goodbye World"}
  var currentIndex := 0

  // Handle button clicks by updating the current message and triggering
  // an update to the expected node on the next frame.
  var toogleTextHandler = func(e Event) {
    if currentIndex == 0 {
      currentIndex = 1
    } else if currentIndex == 1 {
      currentIndex = 0
    }
    e.Target().Invalidate()
  }

  return NanoWindow(NewBuilder(), Title("Test Title"), Width(640), Height(480), FrameRate(24), Children(func(b Builder) {
    Label(b, FlexWidth(1), FlexHeight(1), Text(messages[currentIndex]))
    Button(b, FlexWidth(1), FlexHeight(1), OnClick(toggleTextHandler), Text("Update"))
  }))
}

func main() {
  win, err := createWindow()
  if err != nil {
    panic(err)
  }
  win.(Window).Init()
}
```

# Development environment
This project is being actively developed on OS X and Linux and should build properly in both environments.

If you're on Windows, and interested in contributing, you'll need to get some kind of Posix-ish environment working (probably MingW these days?), and things should probably mostly work? Please let us know if there's anything we can do to help.

### Manual Prerequisites
You'll need to get the following installed on your computer in order to proceed:
* Git
* Make
* Any version of Go (since Go is now bootstrapped)

### Notes
We will download, build and install a specific version of Go and any other tools into the `${PROJECT_PATH}/lib` folder to ensure that all development happens against the same version everywhere.

We are currently building against the [Nanovg](https://github.com/memononen/nanovg) 2d drawing library. I expect this to change in the future in order to deliver  support for rich text rendering. The primary option being considered is [Skia](https://skia.org/), but has not been integrated because the C interface is incomplete and notoriously difficult to rationalize and the build dependencies are also non-trivial. I also spent quite a lot of time getting Cairo working, but ran into difficulties with Pango (rich text layout) and temporarily gave up.

If you have experience with Skia or Cairo/Pango, and would like to contribute, please let us know!

## Download and install
```
git clone https://github.com/waybeams/waybeams.git .
cd golang
make dev-install
```

## Run tests
```
make test
```
Or to get verbose test output:
```
make test-v
```

I use a Python script called, "[when-changed.py](https://github.com/joh/when-changed)" to watch source files and re-run `make test` whenever a file changes. I place the file in my path and use the following command:
```
when-changed.py src/**/*.go  -c "make test"
```

## Build & run binary for development
This should build the binaries from the latest sources on your computer
```
make build
```
Or to build & run in one step
```
make run
make run-anim
```
Build artifacts can be found in `./out`
