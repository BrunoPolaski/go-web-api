package models

import (
	"github.com/BrunoPolaski/go-web-api/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func SelectProducts() []Product {
	db := db.ConnectToDB()
	selectProducts, err := db.Query("select * from products")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func InsertProduct(name, description string, price float64, amount int) {
	db := db.ConnectToDB()
	insertProduct, err := db.Prepare("insert into products(name, description, price, amount) values(?, ?, ?, ?)")

	if err != nil {
		panic(err.Error())
	}

	insertProduct.Exec(name, description, price, amount)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectToDB()

	stmt := "DELETE FROM products WHERE id = ?"

	deleteProduct, err := db.Prepare(stmt)

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}

func SelectProductById(id string) Product {
	db := db.ConnectToDB()

	result := db.QueryRow("SELECT * FROM products WHERE id = ?", id)

	p := Product{}
	err := result.Scan(&p.Id, &p.Name, &p.Description, &p.Price, &p.Amount)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return p
}

func UpdateProduct(id, amount int, name, description string, price float64) {
	db := db.ConnectToDB()

	stmt := "UPDATE products SET name = ?, description = ?, price = ?, amount = ? WHERE id = ?"

	editResult, err := db.Prepare(stmt)

	if err != nil {
		panic(err.Error())
	}

	editResult.Exec(name, description, price, amount, id)

	defer db.Close()
}
