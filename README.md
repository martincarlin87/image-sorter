# Image Sorter



## Build

```bash
go build
```

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
