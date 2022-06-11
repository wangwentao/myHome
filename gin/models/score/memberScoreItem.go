package score

import (
	"myHome/gin/models"
	"time"
)

type MemberScoreItem struct {
	models.PGBaseModel
	RecordId  string
	MemberId  string
	OpDate    time.Time
	ScoreType string
	ScoreName string
	Score     int8
	TypeImage string
	Remark    string
}

func (msi *MemberScoreItem) TableModel() string {

	return "member_score_item"
}
