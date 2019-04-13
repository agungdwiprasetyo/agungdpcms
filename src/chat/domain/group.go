package domain

// Group model
type Group struct {
	ID       int        `gorm:"column:id; primary_key:yes" json:"id"`
	Name     string     `gorm:"column:name" json:"name"`
	Messages []*Message `gorm:"foreignkey:group_id; association_foreignkey:id"`
}

func (g *Group) TableName() string {
	return "chat_groups"
}
