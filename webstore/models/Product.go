package models

import (
	"study.co.jp/webstore/db"
)

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

func CreateProduct(name, description string, price float64, quantity int) {
	connection := db.Connect()
	defer connection.Close()

	//TODO: Check how to prevent injections and others here
	insert, err := connection.Prepare("INSERT INTO PRODUCTS(NAME, DESCRIPTION, PRICE, QUANTITY) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(name, description, price, quantity)
}

func UpdateProduct(product Product) {
	connection := db.Connect()
	defer connection.Close()

	update, err := connection.Prepare("UPDATE PRODUCTS SET NAME=$1, DESCRIPTION=$2, PRICE=$3, QUANTITY=$4 WHERE ID=$5")
	if err != nil {
		panic(err.Error())
	}

	//TODO: All these instructions needs better error handling
	update.Exec(product.Name, product.Description, product.Price, product.Quantity, product.Id)

}

func DeleteProduct(productId string) {
	connection := db.Connect()
	defer connection.Close()

	delete, err := db.Connect().Prepare("DELETE FROM PRODUCTS WHERE ID=$1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(productId)
}

func GetProduct(productId string) Product {
	connection := db.Connect()
	defer connection.Close()

	selectedProducts, err := connection.Query("SELECT * FROM PRODUCTS WHERE ID=$1", productId)
	if err != nil {
		panic(err.Error())
	}

	response := Product{}

	for selectedProducts.Next() {
		p := Product{}

		err := selectedProducts.Scan(&p.Id, &p.Name, &p.Description, &p.Price, &p.Quantity)
		if err != nil {
			panic(err.Error())
		}

		response = p
	}

	return response
}
