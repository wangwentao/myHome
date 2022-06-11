package sys

import "myHome/gin/models"

type TenantType struct {
	models.PGBaseModel
	Type           string
	TypeDesc       string
	MemberType     string
	MemberTypeDesc string
}

func (tp *TenantType) name() string {

	return "tenant_type"
}
