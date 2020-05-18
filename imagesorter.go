package main

import (
	"bytes"
    "log"
)

func ImageSorterTest(buffer *bytes.Buffer, name string) {
    log.Println("ImageSorterTest()")

    buffer.WriteString("my response")
}