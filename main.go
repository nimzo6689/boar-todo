package main

import (
	"flag"
	"fmt"

	"github.com/nimzo6689/boar-todo/config"
)

func main() {
	version := flag.Bool("version", false, "Print version info")
	flag.Parse()

	if *version {
		fmt.Printf("%s v%s\n", config.AppName, config.Version)
		return
	}
}
