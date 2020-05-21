package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var (
	apiHost string
	apiPort string
	debug   bool
)

func init() {
	flag.StringVar(&apiHost, "host", "localhost", "Hostname")
	flag.StringVar(&apiPort, "port", "9090", "Port")
	flag.BoolVar(&debug, "debug", false, "Debug mode")
}

func main() {
	flag.Parse()

	router := mux.NewRouter()

	router.HandleFunc("/", Welcome)
	router.HandleFunc("/submit", Submit).Methods("POST")

	// Assets
	router.HandleFunc("/assets/alpine.js", Alpine)
	router.HandleFunc("/assets/dms2dec.js", Dms2Dec)
	router.HandleFunc("/assets/exif-js.js", Exif)
	router.HandleFunc("/assets/jquery-3.5.1.min.js", JQuery)
	router.HandleFunc("/assets/sweetalert2.min.css", SweetAlertCss)
	router.HandleFunc("/assets/sweetalert2.min.js", SweetAlertJs)
	router.HandleFunc("/assets/tailwind.min.css", Tailwind)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", apiHost, apiPort),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Starting Image Sorter on http://%s:%s\n", apiHost, apiPort)

	url := "http://localhost:9090"

	openUrl(url);

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}

}

func openUrl(url string) error {
    var cmd string
    var args []string

    switch runtime.GOOS {
    	case "windows":
        	cmd = "cmd"
        	args = []string{"/c", "start"}
    	case "darwin":
        	cmd = "open"
    	default: // "linux", "freebsd", "openbsd", "netbsd"
        	cmd = "xdg-open"
    }

    args = append(args, url)
    
    return exec.Command(cmd, args...).Start()
}

func Submit(w http.ResponseWriter, request *http.Request) {

	log.Println("Submit()")

	buffer := new(bytes.Buffer)

	ImageSorterSubmit(buffer, request)

	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("Error")
	}
	
}

func Welcome(writer http.ResponseWriter, request *http.Request) {
	b, _ := Asset("assets/index.html")
	writer.Write(b)
}

// Assets
func Alpine(writer http.ResponseWriter, request *http.Request) {
	b, _ := Asset("assets/alpine.js")
	writer.Write(b)
}

func Dms2Dec(writer http.ResponseWriter, request *http.Request) {
	b, _ := Asset("assets/dms2dec.js")
	writer.Write(b)
}

func Exif(writer http.ResponseWriter, request *http.Request) {
	b, _ := Asset("assets/exif-js.js")
	writer.Write(b)
}

func JQuery(writer http.ResponseWriter, request *http.Request) {
	b, _ := Asset("assets/jquery-3.5.1.min.js")
	writer.Write(b)
}

func SweetAlertCss(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/css")

	b, _ := Asset("assets/sweetalert2.min.css")
	writer.Write(b)
}

func SweetAlertJs(writer http.ResponseWriter, request *http.Request) {
	b, _ := Asset("assets/sweetalert2.min.js")
	writer.Write(b)
}

func Tailwind(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/css")

	b, _ := Asset("assets/tailwind.min.css")
	writer.Write(b)
}

