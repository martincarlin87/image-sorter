# Image Sorter

## Dependencies

```
go get github.com/gorilla/mux
go get -u github.com/go-bindata/go-bindata/...
```

## Build

Update the code on the Windows VM, because there are no user credentials, the easiest way to do that is by:

```
git fetch
git reset --hard origin/master
```

and then

```bash
go-bindata -o assets.go assets
go build
```

## Download

To download the .exe from the build server, open Microsoft Remote Desktop and click the pencil icon on the PC. From there, click `Folders` and choose a directory of your choice to be shared between the remote PC and your local machine.

## Dev

This will regenerate the assets, build the app and launch the executable in debug mode:

```bash
go-bindata -o assets.go assets && go build && ./image-sorter -debug
```

This also allows changes to be made to assets on the fly without having to rebuild after every change.

## Run

```bash
Usage of ./image-sorter:
  -debug
    	Debug mode
  -host string
    	Hostname (default "localhost")
  -port string
    	Port (default "9090")
```

```bash
./image-sorter # defaults to localhost:9090
```

## Assets

To update the assets in `assets.go`, install [go-bindata](https://github.com/go-bindata/go-bindata)

```bash
go get github.com/go-bindata/go-bindata
go-bindata -o assets.go assets
```
