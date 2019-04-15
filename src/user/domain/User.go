package domain

// User domain model
type User struct {
	ID       int    `gorm:"column:id; primary_key:yes" json:"id"`
	Username string `gorm:"column:username; unique_index" json:"username"`
	Password string `gorm:"column:password; unique_index" json:"password,omitempty"`
	Name     string `gorm:"column:name" json:"name,omitempty"`
	Token    string `gorm:"-" json:"token,omitempty"`
	RoleID   int    `gorm:"column:role_id; index" json:"roleId,omitempty"`
	Role     *Role  `gorm:"foreignkey:role_id; association_foreignkey:id" json:"role,omitempty"`
}

// TableName return db table name for user model
func (u *User) TableName() string {
	return "users"
}
