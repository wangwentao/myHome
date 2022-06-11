package sys

import "myHome/gin/models"

type Tenant struct {
	models.PGBaseModel
	TenantId   string
	TenantName string
	TenantType string
}

func (t *Tenant) TableModel() string {

	return "tenant"
}
