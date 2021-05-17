package model

import "time"

type (
	Category struct {
		IdCategory int    `json: "id_category"`
		Name       string `json: "name"`
	}
	Product struct {
		IdProduct  int       `json: "id_product"`
		Name       string    `json: "name"`
		Price      float32   `json: "price"`
		SaleDate   time.Time `json: "sale_date"`
		IdCategory int       `json: "id_category"`
	}

	Products []Product
)
