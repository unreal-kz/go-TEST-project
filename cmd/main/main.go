package main

import (
	"log"
)

func main() {
	data, err := utils.readJSONfile("issues-01.json")
	if err != nil {
		log.Fatal(err)
	}
}
