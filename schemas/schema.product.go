package schemas

type SchemaProduct struct {
	ID         string `json:"id" validate:"uuid"`
	Name       string `json:"name" validate:"required,lowercase"`
	Image      string `json:"image" validate:"required"`
	SKU        uint64 `json:"sku" validate:"required,numeric"`
	Price      uint64 `json:"price" validate:"required,numeric"`
	OutletID   string `json:"outlet_id" validate:"required,uuid"`
	SupplierID string `json:"supplierid" validate:"required,uuid"`
}
