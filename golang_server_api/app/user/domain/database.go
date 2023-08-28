package domain

type MenuModel struct {
	Name string `gorm:"column:name" json:"name"`
}
