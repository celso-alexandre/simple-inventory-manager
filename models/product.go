package models

import "github.com/celso-alexandre/simple-inventory-manager/db"

type Product struct {
	ID             int    `json:"id"`
	Uuid           string `json:"uuid"`
	Name           string `json:"name"`
	Barcode        string `json:"barcode"`
	ProductGroupId int    `json:"productGroupId"`
}

func (p *Product) SaveScan() error {
	res := db.DB.QueryRow(`
		INSERT INTO "Products" ("name", "barcode", "productGroupId") 
		VALUES ($1, $2, $3)
		RETURNING "id", "uuid", "name", "barcode", "productGroupId"
	`, p.Name, p.Barcode, p.ProductGroupId)
	err := res.Scan(&p.ID, &p.Uuid, &p.Name, &p.Barcode, &p.ProductGroupId)
	return err
}
