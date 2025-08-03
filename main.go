package main

import (
	"log"

	"startGo/calculator"
	"startGo/web"
)

func main() {
	cfg := calculator.PriceConfig{BasePrice: 50, FixedCost: 20, PricePerCm: 10, PricePerHour: 150}
	srv, err := web.NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Starting server on :8080")
	if err := srv.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
