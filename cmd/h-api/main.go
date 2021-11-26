package main

import (
	"log"

	"github.com/gentildpinto/h-api/internal/app"
	"github.com/subosito/gotenv"
)

var (
	version = "0.0.1"
)

func init() {
	gotenv.Load()
}

func main() {
	if err := app.Run(version); err != nil {
		log.Fatal(err)
	}
	println("Hello, World!")
}
