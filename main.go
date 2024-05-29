package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nimzo6689/boar-todo/config"
	"github.com/nimzo6689/boar-todo/ui"
)

func main() {
	version := flag.Bool("version", false, "Print version info")
	flag.Parse()

	if *version {
		fmt.Printf("%s v%s\n", config.AppName, config.Version)
		return
	}

	app := ui.NewWindow()
	err := app.Run()
	if err != nil {
		fmt.Printf("Failed to open gui: %v", err)
		os.Exit(1)
	}
}
