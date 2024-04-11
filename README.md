# Simple Inventory Manager
Manage your inventory by scanning with a camera eiter a barcode or qrcode.

## Notes
Migrate: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

## Goals
- [ ] GET /product-scan
  qrcode/barcode
  Can either respond with registered product information or 404
- [ ] PUT /product
  barcode, description, quantity, quantityAction: replace|increment|subtract, categoryId, minimalQty
- [ ] Consider using Htmlx (at least for the desktop report)
