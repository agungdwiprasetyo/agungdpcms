package domain

// Skill model
type Skill struct {
	ID         int     `gorm:"column:id; primary_key:yes" json:"id"`
	ResumeID   int     `gorm:"column:resume_id; index" json:"resumeId"`
	Type       string  `gorm:"column:type" json:"type"`
	Name       string  `gorm:"column:name" json:"name"`
	Percentage float64 `gorm:"column:percentage" json:"percentage"`
}

// TableName for skill
func (s *Skill) TableName() string {
	return "resume_skills"
}
