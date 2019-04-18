package domain

// Achievement model
type Achievement struct {
	ID          int    `gorm:"column:id; primary_key:yes" json:"id"`
	ResumeID    int    `gorm:"column:resume_id; index" json:"resumeId"`
	Name        string `gorm:"column:name" json:"name"`
	Appreciator string `gorm:"column:appreciator" json:"appreciator"`
	Year        int32  `gorm:"column:year" json:"year"`
	Picture     string `gorm:"column:picture" json:"picture"`
}

// TableName Achievement
func (a *Achievement) TableName() string {
	return "resume_achievements"
}
