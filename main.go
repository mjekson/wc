package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	//Defining a boolean flag -l to count lines instead
	lines := flag.Bool("l", false, "Count lines")
	//Defining a boolean flag -b to count Bytes instead
	bytess := flag.Bool("b", false, "Count bytes")
	//Parsing the flags provided by the user
	flag.Parse()

	//Calling
	fmt.Println(count(os.Stdin, *lines, *bytess))
}

func count(r io.Reader, countLines, countBytes bool) int {
	sc := bufio.NewScanner(r)
	if !countLines {
		sc.Split(bufio.ScanWords)
	}
	if !countBytes {
		sc.Split(bufio.ScanBytes)
	}

	wc := 0
	for sc.Scan() {
		wc++
	}
	return wc
}
