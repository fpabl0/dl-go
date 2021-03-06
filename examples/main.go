package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fpabl0/dl-go"
)

type progresser struct{}

func (*progresser) Before() { fmt.Println("Start download") }
func (*progresser) After()  { fmt.Println("\nFinish download") }
func (*progresser) Progress(progress uint64, total uint64) {
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
		&progresser{},
	)

	if err != nil {
		panic(err)
	}

	os.RemoveAll("./archive.zip")
}
