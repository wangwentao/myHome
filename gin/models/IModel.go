package models

type IModel interface {
	TableModel() string
	RefreshTime()
}
