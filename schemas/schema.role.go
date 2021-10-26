package schemas

type SchemaRole struct {
	ID         string   `json:"id" validate:"uuid"`
	RoleName   string   `json:"role_name" validate:"required,lowercase"`
	RoleAccess []string `json:"role_access" validate:"required"`
}
