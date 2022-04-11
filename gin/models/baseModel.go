package models

import (
	"time"
)

type PGBaseModel struct {
	CreateTime time.Time `gorm:"<-:create"`
	UpdateTime time.Time
}

func (bm *PGBaseModel) RefreshTime() {

	if bm.CreateTime.IsZero() {
		bm.CreateTime = time.Now()
	}
	if bm.UpdateTime.IsZero() {
		bm.UpdateTime = time.Now()
	}
}
