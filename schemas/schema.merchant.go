package schemas

type SchemaMerchant struct {
	ID         string `json:"id" validate:"required,uuid"`
	Name       string `json:"name" validate:"required,lowercase" schema:"name"`
	Phone      string `json:"phone" validate:"required,numeric" schema:"phone"`
	Address    string `json:"address" validate:"required,max=1000" schema:"address"`
	Logo       string `json:"logo" schema:"logo"`
	SupplierID string `json:"supplier_id" validate:"required,uuid" schema:"supplier_id"`
}
