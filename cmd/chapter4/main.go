package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("missing args")
	}

	dir := args[0]
	typ := args[1]

	g, err := NewGenerator(dir, typ)
	if err != nil {
		log.Fatal(err)
	}

	src, err := g.Run()
	if err != nil {
		log.Fatal(err)
	}

	// ファイルに書き出す
	path := filepath.Join(dir, fmt.Sprintf("%s_gen.go", strings.ToLower(typ)))
	err = ioutil.WriteFile(path, src, 0644)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
