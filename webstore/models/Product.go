package models

import "study.co.jp/webstore/db"

type Product struct {
	Name, Description string
	Price             float64
	Id, Quantity      int
}

func GetProducts() []Product {
	connection := db.Connect()
	defer connection.Close()

	selectProducts, err := connection.Query("SELECT * FROM PRODUCTS")
	if err != nil {
		panic(err.Error())
	}

	products := []Product{}

	for selectProducts.Next() {
		p := Product{}

		err := selectProducts.Scan(&p.Id, &p.Name, &p.Description, &p.Price, &p.Quantity)
		if err != nil {
			panic(err.Error())
		}

		products = append(products, p)
	}
	return products
}
