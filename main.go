package main

import (
	"log"
	"github.com/TiiQu-Network/claim-verifier-prototype/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
