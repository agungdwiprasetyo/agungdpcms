package domain

// Language model
type Language struct {
	ID   int    `gorm:"column:id; primary_key:yes" json:"id"`
	Type string `gorm:"column:type" json:"type"`
	Name string `gorm:"column:name" json:"name"`
	Icon string `gorm:"column:icon" json:"icon"`
}

// TableName for store language master
func (l *Language) TableName() string {
	return "master_languages"
}
