// Hat is cat for human.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"golang.org/x/term"
)

var nFlag = flag.Bool("n", false, "show line number")

// Main prints each file in the argument list to the standard output.
func main() {
	flag.Parse()

	fileNames := flag.Args()

	if len(fileNames) == 0 {
		dump(os.Stdin)
		return
	}

	for _, fname := range fileNames {
		f, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		defer f.Close()
		dump(f)
	}
}

// Dump prints the content of f to stdout.
func dump(f *os.File) {
	// Print header with filename.
	width, _, err := term.GetSize(0)
	if err != nil {
		return
	}
	for i := 0; i < width; i++ {
		fmt.Printf("%c", '-')
	}
	fmt.Printf("FILE: %s\n", f.Name())
	for i := 0; i < width; i++ {
		fmt.Printf("%c", '-')
	}

	// Print file content.
	s := bufio.NewScanner(f)
	lineNo := 1
	for s.Scan() {
		if *nFlag {
			fmt.Printf("%6d\t", lineNo)
		}
		fmt.Printf("%s\n", s.Text())
		lineNo++
	}
}
