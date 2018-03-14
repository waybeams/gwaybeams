# Getting Started

Step one: Install some version of Go to the system > 1.4. This version will be used purely to build our local version.

Once this is done, change to this directory and run the following (on Darwin/OS X or Linux):

```
make dev-install
source setup-env.sh
```
=======

# Development environment

## Manual Prerequisites
You'll need to get the following installed on your computer in order to proceed:
* Git
* Make
* Some version of Go (since it's now bootstrapped)
* _Probably some Python (2.7?)i or other (3.x?) in order to build Skia_

*Notes:*
We will download, build and install a specific version of Go and any other tools into the `${PROJECT_PATH}/lib` folder to ensure that all development happens against the same source tree.

We are currently building against Skia master (trunk). This is almost definitely not desirable, but I'm still too unfailiar with Skia to make a clear decision as to which branch to build against. I expect this project (like Skia) to be deployed into a relatively large number of operating system environments.

## Download and install
```
git clone https://github.com/lukebayes/findingyou.git .
cd golang
make dev-install
# Wait and wait for skia.so to get built
```

## On Linux

I was unable to build Skia with an error of:
```
GL/glx.h: No such file or directory
```

This was fixed with the following command:
```
sudo apt-get install libglu1-mesa-dev freeglut3-dev mesa-common-dev
```

## Run tests
```
make test
```

I use a Python script called, "when-changed.py" (search Google) to watch source files and re-run `make test` whenever a file changes. This seems to work just fine.


## Build & run binary for development
This should build the binary from latest sources on your system and 
```
make run
```

