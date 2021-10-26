package schemas

type SchemaCustomer struct {
	ID      string `json:"id" validate:"uuid"`
	Name    string `json:"name" validate:"required,lowercase"`
	Phone   string `json:"phone" validate:"required,gte=12"`
	Address string `json:"address" validate:"required,max=1000"`
}
