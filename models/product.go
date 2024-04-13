package models

import (
	"database/sql"
	"errors"

	"github.com/celso-alexandre/simple-inventory-manager/db"
)

type Product struct {
	Id              int64  `json:"id"`
	Uuid            string `json:"uuid"`
	Name            string `json:"name"`
	Barcode         string `json:"barcode"`
	ProductGroupId  int    `json:"productGroupId"`
	UpdatedByUserId int64  `json:"updatedByUserId"`
}

func (p *Product) SaveScan() error {
	var res *sql.Row
	if p.Id > 0 {
		res = db.DB.QueryRow(`
			SELECT "id", "uuid", "name", "barcode", "productGroupId"
			FROM "Products"
			WHERE "id" = $1
		`, p.Id)
		return res.Scan(&p.Id, &p.Uuid, &p.Name, &p.Barcode, &p.ProductGroupId)
	}
	if p.Barcode != "" {
		res = db.DB.QueryRow(`
			INSERT INTO "Products" ("barcode", "updatedByUserId")
			VALUES ($1, $2)
			ON CONFLICT DO UPDATE SET "barcode" = $1
			RETURNING "id", "uuid", "name", "barcode", "productGroupId"
		`, p.Barcode, p.UpdatedByUserId)
		return res.Scan(&p.Id, &p.Uuid, &p.Name, &p.Barcode, &p.ProductGroupId)
	}
	return errors.New("id or barcode is required")
}

func (p *Product) Update() error {
	res := db.DB.QueryRow(`
		UPDATE "Products"
		SET "name"            = $1,
			 "barcode"         = $2,
			 "productGroupId"  = $3,
			 "updatedByUserId" = $5,
			 "updatedAt"       = now()
		WHERE "id" = $4
		RETURNING "id", "uuid", "name", "barcode", "productGroupId"
	`, p.Name, p.Barcode, p.ProductGroupId, p.Id, p.UpdatedByUserId)
	return res.Scan(&p.Id, &p.Uuid, &p.Name, &p.Barcode, &p.ProductGroupId, &p.UpdatedByUserId)
}
