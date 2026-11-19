package main

import (
	"embed"
	"fmt"
	"io/fs"
	"strings"
)

//go:embed testdata/embed/hello.txt
var s string

var (
	//go:embed testdata/embed/hello.txt
	b []byte

	//go:embed testdata/embed/hello.txt testdata/embed/sub/*.txt
	//go:embed testdata/embed/sub
	dirFS embed.FS

	//go:embed all:testdata/embed/sub
	allFS embed.FS
)

func main() {
	var _ fs.FS = dirFS
	var _ fs.ReadFileFS = dirFS
	var _ fs.ReadDirFS = dirFS

	entries, err := dirFS.ReadDir("testdata/embed/sub")
	if err != nil {
		panic(err)
	}
	var names []string
	for _, e := range entries {
		names = append(names, e.Name())
	}

	allEntries, err := allFS.ReadDir("testdata/embed/sub")
	if err != nil {
		panic(err)
	}
	var allNames []string
	for _, e := range allEntries {
		allNames = append(allNames, e.Name())
	}

	f, err := dirFS.Open("testdata/embed/sub")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, ok := f.(fs.ReadDirFile); !ok {
		panic("not fs.ReadDirFile")
	}

	b1, err := dirFS.ReadFile("testdata/embed/hello.txt")
	if err != nil {
		panic(err)
	}
	b1[0] = 'H'
	b2, err := dirFS.ReadFile("testdata/embed/hello.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(strings.TrimSpace(s))
	fmt.Println(strings.TrimSpace(string(b)))
	fmt.Println(strings.TrimSpace(string(b2)))
	fmt.Println(strings.Join(names, ","))
	fmt.Println(strings.Join(allNames, ","))
}

// Output:
// hello from embed
// hello from embed
// hello from embed
// a.txt,b.txt
// .dotfile,_underfile,a.txt,b.txt
