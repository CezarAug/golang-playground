package main

import (
	"fmt"
	"go-crawler/internal/fetcher"
	"go-crawler/internal/models"
	"go-crawler/internal/processor"
	"time"
)

func main() {
	start := time.Now()
	// buffer
	priceChannel := make(chan models.PriceDetail)
	// priceChannel := make(chan models.PriceDetail, 4)
	// len & cap (current buffer and buffer limit)

	done := make(chan bool)

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAVG(priceChannel, done)

	<-done //Stops execution until a true is received here

	fmt.Printf("\n Time: %s", time.Since(start))
}
