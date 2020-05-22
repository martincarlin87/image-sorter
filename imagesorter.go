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


    	// keep images with no GPS Data in an `unsorted` directory
    	if location == "" || location == "null" {
    		location = "unsorted"
    	} else if (ignore) {
    		// image is too far from the structure, sort in to a separate directory
    		location = "outwith_threshold"
    	}

    	fmt.Println("\n---")
		log.Println(image)
		log.Println(location)
		log.Println(line)
		fmt.Println("\n---")

		sortImage(image, location, line)

		i++
	}

	log.Println("\n")
	log.Println("DONE!")

    buffer.WriteString("success")
}

func sortImage(image string, directory string, line string) bool {

	// create the directory if it doesn't exist
	newpath := filepath.Join(".", "output", line, directory)
	os.MkdirAll(newpath, os.ModePerm)

	// https://shapeshed.com/copy-a-file-in-go/
	from, err := os.Open(filepath.Join(".", "input", image))

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

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }

    return false
}
