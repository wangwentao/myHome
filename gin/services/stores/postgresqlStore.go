package stores

import (
	"gorm.io/gorm/clause"
	"myHome/gin/configs"
	"myHome/gin/models"
	"myHome/gin/utils"
)

func NewModel(im models.IModel) error {

	im.RefreshTime()
	db := configs.PGStore.Table(im.TableModel()).Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(im)

	return db.Error
}

func UpdateModel(im models.IModel) error {

	im.RefreshTime()
	err := configs.PGStore.Table(im.TableModel()).Save(im).Error
	utils.CheckErr(err)
	return err
}
