package main

import (
	"fmt"
	"log"

	"github.com/gentildpinto/h-api/internal/app"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	fmt.Println("Starting server...")
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
