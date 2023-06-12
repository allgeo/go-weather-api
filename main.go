package main

import (
	"blog-apis/api"
	"flag"
)

func main() {
	flag.Parse()
	api.Run()
}
