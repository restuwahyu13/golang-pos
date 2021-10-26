package schemas

type SchemaOutlet struct {
	ID        string `json:"id" validate:"required,uuid"`
	Name      string `json:"name" validate:"required,lowercase"`
	Phone     string `json:"phone" validate:"required,gte=12"`
	Address   string `json:"address" validate:"required,max=1000"`
	MerchatID string `json:"merchant_id" validate:"required,uuid"`
}
