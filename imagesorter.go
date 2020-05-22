package main

import (
	"bytes"
	"fmt"
	"io"
    "log"
    "net/http"
    "path/filepath"
    "os"
    "strconv"
)


func ImageSorterSubmit(buffer *bytes.Buffer, request *http.Request) {
    log.Println("ImageSorterSubmit()")

    // remove unsorted directory
    unsortedpath := filepath.Join(".", "output", "unsorted")
    os.RemoveAll(unsortedpath)

    thresholdpath := filepath.Join(".", "output", "outwith_threshold")
    os.RemoveAll(thresholdpath)

    request.ParseMultipartForm(1024 * 1024 * 16)

    i := 0

    // empty slices to keep track of the lines we have seen so that we can delete the directory the first time we see it before populating with new data
    lines := []string{}

    // form data is in a slice, which is unordered
    // divide by 4 because there are four arrays of (hopefully) equal length, one for images, one for locations, one for lines and one for if the image should be ignored
    for i < len(request.Form) / 4 {
    	image := request.FormValue("image[" + strconv.Itoa(i) + "]")
    	location := request.FormValue("location[" + strconv.Itoa(i) + "]")
    	line := request.FormValue("line[" + strconv.Itoa(i) + "]")
    	ignore, err := strconv.ParseBool(request.FormValue("ignore[" + strconv.Itoa(i) + "]"))

	    if err != nil {
	        ignore = false
	    }

    	// remove `output/{line}` directory
    	if (!contains(lines, line)) {
    		linepath := filepath.Join(".", "output", line)
    		os.RemoveAll(linepath)

    		// add `line` to lines slice
    		lines = append(lines, line)
    	}

    	path := ""

    	// keep images with no GPS Data in an `unsorted` directory
    	if location == "" || location == "null" {
    		path = unsortedpath
    	} else if (ignore) {
    		// image is too far from the structure, sort in to a separate directory
    		path = thresholdpath
    	} else {
    		path = filepath.Join(".", "output", line, location)
    	}

    	fmt.Println("\n")
		log.Println(image)
		log.Println(location)
		log.Println(line)
		fmt.Println("\n")

		sortImage(image, path)

		i++
	}

	fmt.Println("\n")
	log.Println("DONE!")

    buffer.WriteString("success")
}

func sortImage(image string, path string) bool {

	// create the directory if it doesn't exist
	os.MkdirAll(path, os.ModePerm)

	// https://shapeshed.com/copy-a-file-in-go/
	from, err := os.Open(filepath.Join(".", "input", image))

	if err != nil {
	    log.Fatal(err)
	}
	
	defer from.Close()

	to, err := os.OpenFile(path + "/" + image, os.O_RDWR|os.O_CREATE, 0666)
	
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

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }

    return false
}
