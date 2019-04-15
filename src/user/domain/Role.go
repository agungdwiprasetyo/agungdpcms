package domain

// Role domain model
type Role struct {
	ID   int    `gorm:"column:id; primary_key:yes" json:"id"`
	Slug string `gorm:"column:slug; unique_index" json:"slug,omitempty"`
	Name string `gorm:"column:name" json:"name,omitempty"`
}

// TableName return table name of Role domain
func (r *Role) TableName() string {
	return "roles"
}
