package dl

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownloadInvalidDest(t *testing.T) {
	err := Download(
		"https://github.com/briandowns/spinner/archive/master.zip",
		"temp/archive.zip",
		nil,
	)
	if assert.NotNil(t, err) {
		assert.EqualValues(t, err.Error(), "open temp/archive.zip: no such file or directory")
	}
}

func TestDownloadInvalidURL(t *testing.T) {
	err := Download(
		"https://github1111.com/briandowns/spinner/archive/master.zip",
		"./archive.zip",
		nil,
	)
	if assert.NotNil(t, err) {
		assert.EqualValues(t, err.Error(), "Get https://github1111.com/briandowns/spinner/archive/master.zip: dial tcp: lookup github1111.com: no such host")
	}
}

func TestDownloadNotFound(t *testing.T) {
	err := Download(
		"https://github.com/briandowns/spinner2/archive/master.zip",
		"./archive.zip",
		nil,
	)
	if assert.NotNil(t, err) {
		assert.EqualValues(t, err.Error(), "resource not found")
	}
}

type mockProgresser struct{}

func (*mockProgresser) Before() {}
func (*mockProgresser) After()  {}
func (*mockProgresser) Progress(progress uint64, total uint64) {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	if total == 0 {
		fmt.Printf("\rDownloading... %d complete", progress)
	} else {
		fmt.Printf("\rDownloading... %d/%d complete", progress, total)
	}
}

func TestDownloadNoError(t *testing.T) {

	err := os.Mkdir("./temp", os.ModePerm)
	if !assert.Nil(t, err) {
		return
	}
	defer os.RemoveAll("./temp")

	err = Download(
		"https://github.com/briandowns/spinner/archive/master.zip",
		"./temp/archive.zip",
		&mockProgresser{},
	)
	assert.Nil(t, err)
	f, err := os.Stat("./temp/archive.zip")
	assert.Nil(t, err)
	if assert.NotNil(t, f) {
		assert.False(t, f.IsDir())
	}

}
