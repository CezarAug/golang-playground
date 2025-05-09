package processor

import (
	"fmt"
	"go-crawler/internal/models"
)

// Attention at the <- operator for each channel, one for only reading, and another to write
func ShowPriceAVG(priceChannel <-chan models.PriceDetail, done chan<- bool) {
	var totalPrice float64
	countPrices := 0.0
	for price := range priceChannel {
		totalPrice += price.Value
		countPrices++
		avgPrice := totalPrice / countPrices
		fmt.Printf(
			"[%s]Store: %s | $ %.2f | AVG: %.2f \n",
			price.Timestamp.Format("02-01-2006 15:04:05"),
			price.StoreName, price.Value, avgPrice)
	}

	done <- true
}
