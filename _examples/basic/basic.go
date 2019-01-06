package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/dsbezerra/go-ingresso/ingresso"
)

func main() {
	// Instantiate the Ingresso lib
	ing := ingresso.New(
		ingresso.Partnership("YOUR_PARTNERSHIP_CODE"),
	)

	// Get information about theater 372 - Cinemark BH Shopping
	theater, err := ing.GetTheater("372")
	if err != nil {
		log.Fatal(err)
	}

	// Print out the JSON
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(theater)
}
