package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	MD_SUFFIX = ".md"
)

var (
	recursive = flag.Bool("r", false, "recursively traverse directories")
	pattern   = flag.String("p", "*.md", "pattern to process. use with -r.")
)

func do(filename string) {
	if strings.HasSuffix(filename, MD_SUFFIX) == false {
		log.Printf("%q does not appear to be markdown\n", filename)
		return
	}

	stat, err := os.Stat(filename)
	if err != nil {
		log.Printf("stat %s: %v\n", filename, err)
		return
	}

	if stat.Mode().IsDir() {
		log.Println("directories not yet supported")
		return
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("read %s: %v\n", filename, err)
		return
	}

	processed, err := Process(file)
	if err != nil {
		log.Printf("markdown error in %s: %v\n", filename, err)
		return
	}

	newfile := strings.TrimSuffix(filename, MD_SUFFIX)
	err = ioutil.WriteFile(newfile, processed, stat.Mode())
	if err != nil {
		log.Printf("write %s result to %s: %v\n", filename, newfile, err)
		return
	}
}

func main() {
	flag.Parse()

	// TODO Switch on arguments.

	for _, filename := range flag.Args() {
		do(filename)
	}
}
