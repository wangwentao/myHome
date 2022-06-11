package score

import "myHome/gin/models"

type MemberScore struct {
	models.PGBaseModel
	MemberId   string
	MemberName string
	Score      int8
}

func (ms *MemberScore) TableModel() string {

	return "member_score"
}
