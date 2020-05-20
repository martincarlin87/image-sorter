package main

import (
	"bytes"
	"io"
    "log"
    "net/http"
    "path/filepath"
    "os"
    "strconv"
)


func ImageSorterSubmit(buffer *bytes.Buffer, request *http.Request) {
    log.Println("ImageSorterSubmit()")

    // remove`output` directory
	os.RemoveAll("./output")

    request.ParseMultipartForm(1024 * 1024 * 16);

    i := 0

    // form data is in a slice, which is unordered
    // divide by 2 because there are two arrays of (hopefully) equal length, one for images, one for locations
    for i < len(request.Form) / 2 {
    	image := request.FormValue("image[" + strconv.Itoa(i) + "]")
    	location := request.FormValue("location[" + strconv.Itoa(i) + "]")

		log.Println(image);
		log.Println(location);

		sortImage(image, location)

		i++
	}

	log.Println("TEST")
    buffer.WriteString("my response")
}

func sortImage(image string, directory string) bool {

	// create the directory if it doesn't exist
	newpath := filepath.Join(".", "output", directory)
	os.MkdirAll(newpath, os.ModePerm)

	// https://shapeshed.com/copy-a-file-in-go/
	from, err := os.Open("./" + image)

	if err != nil {
	    log.Fatal(err)
	}
	
	defer from.Close()

	to, err := os.OpenFile(newpath + "/" + image, os.O_RDWR|os.O_CREATE, 0666)
	
	if err != nil {
	    log.Fatal(err)
	}
	
	defer to.Close()

	_, err = io.Copy(to, from)
	
	if err != nil {
	    log.Fatal(err)
	}

	return true

}