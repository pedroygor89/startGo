package main

import (
	"fmt"
	"time"

	"startGo/calculator"
	"startGo/finance"
	"startGo/inventory"
	"startGo/scheduler"
	"startGo/user"
)

func main() {
	cfg := calculator.PriceConfig{BasePrice: 50, FixedCost: 20, PricePerCm: 10, PricePerHour: 150}
	price := calculator.EstimateByDetails(cfg, 1.2, 1.1, 1.0, 15)
	fmt.Printf("Estimated price: %.2f\n", price)

	repo := scheduler.Repository{}
	repo.Add(scheduler.Appointment{
		ID: "1", Client: "Alice", Artist: "Bob",
		Time: time.Now(), Description: "Rose on arm",
	})
	appointments := repo.ListByDate(time.Now())
	fmt.Printf("Appointments today: %d\n", len(appointments))

	ledger := finance.Ledger{}
	ledger.Add(finance.Record{
		Date: time.Now(), Artist: "Bob", Client: "Alice",
		Value: price, Method: "PIX",
	})
	fmt.Printf("Finance records: %d\n", len(ledger.All()))

	store := inventory.Store{}
	store.Add(inventory.Item{Name: "Ink", Quantity: 10, MinAlert: 5})
	store.AdjustQuantity("Ink", -2)
	fmt.Printf("Ink quantity: %d\n", store.Items()[0].Quantity)

	_ = user.User{Name: "Bob", Role: "Tattoo Artist", Phone: "555-1234", Instagram: "@bobtattoo", Email: "bob@example.com"}
}
