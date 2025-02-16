package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/tidwall/gjson"
)

var count = 100000

func main() {
	// Generate a UUID
	for i := 0; i < count; i++ {
		id := uuid.New()
		log.Printf("Generated UUID: %s", id.String())
	}

	// Parse a JSON string
	for i := 0; i < count; i++ {
		json := `{"name": "GitHub Copilot", "age": 1}`
		name := gjson.Get(json, "name")
		log.Printf("Parsed name from JSON: %s", name.String())
	}

	log.Print("Done")
}
