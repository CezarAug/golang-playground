package main

import (
	"fmt"
	"stock/internal/models"
	"stock/internal/services"
)

func main() {
	fmt.Println("Stock")

	stock := services.NewStock()
	items := []models.Item{
		{
			ID:       1,
			Name:     "Phone",
			Quantity: 5,
			Price:    10.82,
		},
		{
			ID:       2,
			Name:     "Shirt",
			Quantity: 5,
			Price:    1000,
		},
		{
			ID:       3,
			Name:     "Mouse",
			Quantity: 5,
			Price:    100.99,
		},
	}

	for _, item := range items {
		err := stock.AddItem(item, "User01")

		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	// fmt.Println(stock)

	// fmt.Println(stock.ListItems())
	for _, item := range stock.ListItems() {
		fmt.Printf("\nID: %d | Name: %s | Quantity: %d | Price: %.2f", item.ID, item.Name, item.Quantity, item.Price)
	}

	fmt.Println("\n============================")
	fmt.Println("Total stock value: ", stock.CalculateTotalCost())

	fmt.Println(stock.RemoveItem(models.Item{ID: 2, Quantity: 1}, "userDelete"))

	for _, item := range stock.ListItems() {
		fmt.Printf("\nID: %d | Name: %s | Quantity: %d | Price: %.2f", item.ID, item.Name, item.Quantity, item.Price)
	}

	// logs := stock.ViewAuditLogs()

	// for _, log := range logs {
	// 	fmt.Printf("\n[%s] Ação: %s - Usuário: %s - Item ID: %d - Quantidade: %d - Motivo: %s",
	// 		log.Timestamp.Format("01/02 15:04:05"), log.Action, log.User, log.ItemId, log.Quantity, log.Reason)
	// }

	itemByName, err := services.Find(items, func(item models.Item) bool {
		return item.Price > 100
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(itemByName)

	test := services.Supplier{
		Number:  "10102020-1",
		Contact: "810707070070070707",
		City:    "Tokyo",
	}

	fmt.Println(test.GetInfo())
	if test.CheckAvailability(10, 15) {
		fmt.Println("Available")
	} else {
		fmt.Println("Not available")
	}
}
