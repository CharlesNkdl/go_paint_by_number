package main

import (
	"log"

	"github.com/charlesNkdl/go_paint_by_number/internal"
)

func main() {
	err := internal.Run()
	if err != nil {
		log.Fatalf("%s", err)
	}
	return
}
