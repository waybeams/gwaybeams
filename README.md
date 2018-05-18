
## This software is experimental. Contributions are welcome, but production use is discouraged

# Waybeams

With Waybeams you can quickly and reliably create and test tiny (<4MB). delightful applications that can thrive on a multitude of surfaces (Windows, macOS, Linux, Android, iOS, Raspberry Pi, Beaglebone, etc.).

According to Merriam Webster, A Waybeam is, ": a beam supporting a way; specifically : either of two longitudinal beams resting on transverse girders and supporting the rails of a road crossing a bridge"

We like to think of Waybeams (the tools here) as providing a solid structural foundation that makes it possible for us to safely and quickly transport enormous quantities of user facing features to production.

![Waybeams Image](media/waybeams-home.jpg)

_[Image](https://www.flickr.com/photos/charlyamato/13417543435/) provided courtesy of [Carlos Amato](https://www.flickr.com/photos/charlyamato/) and the [Creative Commons](https://creativecommons.org/licenses/by-nc-nd/2.0/) license(s)._

Waybeams is built using the [Go language](https://golang.org/) and an OpenGL rendering surface (currently, [NanoVGO](https://github.com/shibukawa/nanovgo))

Waybeams provides (or for now, aspires to provide):
* Simple, fast, reactive and composable GUI toolkit
* Cross-platform, GPU accelerated drawing surface
* Pure Go component declaration and configuration
* Tiny, blazing fast, constraints-based layouts
* Component Trait assignment using web-like selectors
* Headless (insanely fast) environment for UI tests
* Isolated visual environment for test-driven development on UI components
* Automated image rendering surface (from tests) for release validation


