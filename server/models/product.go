package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/celso-alexandre/simple-inventory-manager/server/db"
)

type Product struct {
	Id              int64     `json:"id"`
	Uuid            string    `json:"uuid"`
	Name            string    `json:"name"`
	Barcode         string    `json:"barcode"`
	ProductGroupId  int       `json:"productGroupId"`
	UpdatedByUserId int64     `json:"updatedByUserId"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

func (p *Product) SaveScan() error {
	var res *sql.Row
	if p.Uuid != "" {
		res = db.DB.QueryRow(`
			SELECT "id", "uuid", coalesce("name", '') AS "name", "barcode", "productGroupId", "createdAt", "updatedAt", "updatedByUserId"
			FROM "Products"
			WHERE "uuid" = $1
		`, p.Uuid)
		return res.Scan(&p.Id, &p.Uuid, &p.Name, &p.Barcode, &p.ProductGroupId, &p.CreatedAt, &p.UpdatedAt, &p.UpdatedByUserId)
	} else if p.Barcode != "" {
		res = db.DB.QueryRow(`
			INSERT INTO "Products" ("barcode", "updatedByUserId", "productGroupId")
			VALUES ($1::text, $2, 1)
			ON CONFLICT ("barcode") DO UPDATE SET "barcode" = $1
			RETURNING "id", "uuid", coalesce("name", '') AS "name", "barcode", "productGroupId", "createdAt", "updatedAt", "updatedByUserId"
		`, p.Barcode, p.UpdatedByUserId)
		return res.Scan(&p.Id, &p.Uuid, &p.Name, &p.Barcode, &p.ProductGroupId, &p.CreatedAt, &p.UpdatedAt, &p.UpdatedByUserId)
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
		RETURNING "id", "uuid", coalesce("name", '') AS "name", "barcode", "productGroupId", "createdAt", "updatedAt", "updatedByUserId"
	`, p.Name, p.Barcode, p.ProductGroupId, p.Id, p.UpdatedByUserId)
	return res.Scan(&p.Id, &p.Uuid, &p.Name, &p.Barcode, &p.ProductGroupId, &p.CreatedAt, &p.UpdatedAt, &p.UpdatedByUserId)
}
