package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fpabl0/dl-go"
)

func printProgress(progress uint64, total uint64) {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	if total == 0 {
		fmt.Printf("\rDownloading... %d complete", progress)
	} else {
		fmt.Printf("\rDownloading... %d/%d complete", progress, total)
	}
}

func main() {
	err := dl.Download(
		"https://github.com/briandowns/spinner/archive/master.zip",
		"./archive.zip",
		printProgress,
	)

	if err != nil {
		panic(err)
	}

	os.RemoveAll("./archive.zip")
}
