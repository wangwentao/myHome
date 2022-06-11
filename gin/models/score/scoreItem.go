package score

import "myHome/gin/models"

type ScoreItem struct {
	models.PGBaseModel
	Id    string
	Type  string
	Code  string
	Name  string
	Score int8
	Image string
}

func (si *ScoreItem) TableModel() string {

	return "score_item"
}
