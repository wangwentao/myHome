package sys

import "myHome/gin/models"

type Member struct {
	models.PGBaseModel
	MemberId       string
	TenantId       string
	MemberName     string
	MemberType     string
	MemberTypeDesc string
	OpenId         string
}

func (m *Member) TableModel() string {

	return "member"
}
