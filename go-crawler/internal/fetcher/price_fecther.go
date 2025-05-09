package fetcher

import (
	"go-crawler/internal/models"
	"math/rand"
	"sync"
	"time"
)

// Simulating crawling prices from different websites, and the time it takes to retrieve the information

func FetchPriceFromSite1() models.PriceDetail {
	time.Sleep(1 * time.Second)
	return models.PriceDetail{
		StoreName: "Store1",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchPriceFromSite2() models.PriceDetail {
	time.Sleep(3 * time.Second)
	return models.PriceDetail{
		StoreName: "Store2",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchPriceFromSite3() models.PriceDetail {
	time.Sleep(2 * time.Second)
	return models.PriceDetail{
		StoreName: "Store3",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchAndSendMultiplePrices(priceChannel chan<- models.PriceDetail) {
	time.Sleep(3 * time.Second)
	prices := []float64{
		rand.Float64() * 100,
		rand.Float64() * 100,
		rand.Float64() * 100,
	}

	for _, price := range prices {
		priceChannel <- models.PriceDetail{
			StoreName: "MultiStore",
			Value:     price,
			Timestamp: time.Now(),
		}
	}
}

func FetchPrices(priceChannel chan<- models.PriceDetail) {
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite1()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite2()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite3()
	}()

	go func() {
		defer wg.Done()
		FetchAndSendMultiplePrices(priceChannel)
	}()

	wg.Wait()
	close(priceChannel)

}
