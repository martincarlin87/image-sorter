package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
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
	router.HandleFunc("/test", Test)
	// router.HandleFunc("/{name}", Test)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", apiHost, apiPort),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Starting Image Sorter on http://%s:%s\n", apiHost, apiPort)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func Test(w http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	log.Println("Test()")
	log.Println(vars["name"])
	buffer := new(bytes.Buffer)

	ImageSorterTest(buffer, vars["name"])

	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("Error")
	}
	
}

func Welcome(writer http.ResponseWriter, request *http.Request) {
	b, _ := Asset("assets/index.html")
	writer.Write(b)
}
