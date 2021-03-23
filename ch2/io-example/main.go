package main

import (
	"fmt"
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

	// Create buffer to hold input/output
	input := make([]byte, 4096)

	// Use reader to read input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	// Use writer to write output
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)
}
