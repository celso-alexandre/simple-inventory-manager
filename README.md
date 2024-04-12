# Simple Inventory Manager
Manage your inventory by scanning with a camera eiter a barcode or qrcode.

## Notes
Migrate: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

## Goals
- [X] POST /products-scan
  qrcode/barcode
- [/] PUT /products/:id
  barcode, description, categoryId, minimalQty
  todo: [id, locationId, quantityToAdd]
- [ ] Consider using Htmx (at least for the desktop report)
