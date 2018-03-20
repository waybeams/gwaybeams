
# Epiphyte
===========

Epiphyte helps you use the Go language to quickly, reliably and delightfully build (minuscule to giant) screaming fast graphical applications that can live on ALL the things (desktop, mobile, tablet, TV, Pi, embedded, etc.).

With Epiphyte you can quickly and reliably create tiny, fast and delightful applications that (like the plants) can thrive on top of just about any environment (Windows, OS X, Linux, Android, iOS, Raspberry Pi, Beaglebone, etc.).

<img src="media/epiphyte.png" style="float:right;" />
![](media/epiphyte.jpg)
*Image From, "Botany for young people and common schools" 1868 by Asa Gray [Source](https://commons.wikimedia.org/wiki/File:Botany_for_young_people_and_common_schools_(1868)_(20219036949).jpg)*

Epiphyte is built using the [Go language](https://golang.org/) and [Cairo graphics engine](https://cairographics.org/).

# Getting Started

Step one: Install some version of Go to the system > 1.4. This version will be used purely to build our local version.

Once this is done, change to this directory and run the following (on Darwin/OS X or Linux):

```
source setup-env.sh
make dev-install
```
=======

# Development environment

## Manual Prerequisites
You'll need to get the following installed on your computer in order to proceed:
* Git
* Make
* Some version of Go (since it's now bootstrapped)
* Possibly some dev headers for Cairo

*Notes:*
We will download, build and install a specific version of Go and any other tools into the `${PROJECT_PATH}/lib` folder to ensure that all development happens against the same source tree.

We are currently integrated with the Cairo 2d drawing library. This library was selected over Skia simply because Skia's C interface is still experimental and does not support most of the features we need. Cairo also seems to build more easily.

## Download and install
```
git clone https://github.com/lukebayes/findingyou.git .
cd golang
make dev-install
# Wait and wait for Cairo to be built
```

## On Linux

I'm still having trouble getting this to work on Ubuntu...

## Run tests
```
make test
```
Or to get verbose test output:
```
make test-v
```

I use a Python script called, "when-changed.py" (search Google) to watch source files and re-run `make test` whenever a file changes. This seems to work just fine for me.

## Build & run binary for development
This should build the binaries from the latest sources on your computer
```
make build
```
Or to build & run in one step
```
make run
```
Build artifacts can be found in `./out`
