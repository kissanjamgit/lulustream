package main

import (
	"os"

	"github.com/kissanjamgit/lulustream"
	"resty.dev/v3"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("len(args) < 2")
	}
	lulu := lulustream.New(args[1])
	cr, err := lulu.Resource(resty.New())
	if err != nil {
		panic(err)
	}
	lulu.Download(cr)
}
