# Unofficial Go for Windows

## About

* Go 1.10 is the last release that will run on Windows XP or Windows Vista.
* Go 1.11 requires Windows 7 or later.

This repo helps runing latest Go on those ancient Windows.

## How to use

If you're cross compiling for Windows under Linux/macOS, just apply the patch and you're good to go.

For native build on Windows, you have two choices:
1. download pre-built Go distribution containing necessary patches
   (TODO)
2. build from source
   1. apply patch from this repo
   2. consult https://golang.org/doc/install/source

## Patch

### for maintained releases

* [patch for Go 1.15.1](https://github.com/zhangyoufu/go/compare/go1.15.1...windows.go1.15.diff)
* [patch for Go 1.14.8](https://github.com/zhangyoufu/go/compare/go1.14.8...windows.go1.14.diff)

### for archived releases

* [patch for Go 1.13.15](https://github.com/zhangyoufu/go/compare/go1.13.15...windows.go1.13.diff)
* [patch for Go 1.12.17](https://github.com/zhangyoufu/go/compare/go1.12.17...windows.go1.12.diff)
* [patch for Go 1.11.13](https://github.com/zhangyoufu/go/compare/go1.11.13...windows.go1.11.diff)

## Caveats

* DO NOT expect 3rd-party packages, for example github.com/microsoft/go-winio, to work seamlessly on ancient Windows. They are typically not supported/tested. You are on your own.
* os.UserCacheDir() on NT 5.x returns %AppData% instead of return an error saying that %LocalAppData% is not defined. Applies to Go 1.11 and later.
* os.Pipe() creates named pipe with FILE_FLAG_OVERLAPPED on NT 5.x/6.0. Applies to Go 1.13 and later.
* On NT 5.x/6.0, if an os.File is created with a pipe handle capable of overlapped I/O, the first call to Read/Write will put that handle into netpoll and associated with netpoll's IOCP. This may introduce unexpected result if you're rolling your own asynchronous I/O. Applies to Go 1.13 and later.
* On NT 5.x/6.0, if an os.File is created with a pipe handle not capable of overlapped I/O, the first call to Close will be blocked by any in-progress Read/Write, until the other end of the pipe goes away. Applies to Go 1.13 and later.
* OperatingSystemVersion & SubsystemVersion are hard-coded to NT 5.1. Applies to Go 1.13 and later.
* Binary built with race detector can not run on NT 6.0 and earlier, due to PSAPI_VERSION. Applies to Go 1.14 and later.
