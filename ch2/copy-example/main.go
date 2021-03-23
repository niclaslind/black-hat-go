package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FooReader struct{}

func (fooReader *FooReader) Read(p []byte) (int, error) {
	// Read some data from somewhere, anywhere.
	fmt.Print("in > ")
	return os.Stdin.Read(p)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(p []byte) (int, error) {
	// Write data somewhere
	fmt.Print("out > ")
	return os.Stdout.Write(p)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	// Copy all data from stdin to stdout
	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}
