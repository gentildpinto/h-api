package main

import (
	"log"

	"github.com/gentildpinto/h-api/internal/app"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
