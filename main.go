package main

import (
	"github.com/dmarciniak/loge"
	"fmt"
	"os"
	"strconv"
)

func main() {
	output := loge.LogLoader(loadFilenames())

	for entry := <-output; !entry.IsEOF(); entry = <-output {
		fmt.Println(entry.FileName + " " + entry.Date.String() + " " + strconv.Itoa(entry.LineNo) + ": " + entry.Log)
	}
}

func loadFilenames() []string {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("At least one filename is required")
		os.Exit(1)
	}
	return arguments[1:]
}
