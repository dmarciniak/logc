package main

import (
	"github.com/dmarciniak/loge"
	"fmt"
	"os"
)

func main() {
	output := loge.LogLoader(loadFilenames())

	for entry := <-output; !entry.IsEOF(); entry = <-output {
		fmt.Println(entry.Log)
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
