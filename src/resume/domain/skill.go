package domain

// Skill model
type Skill struct {
	ID          int     `gorm:"column:id; primary_key:yes" json:"id,omitempty"`
	ResumeID    int     `gorm:"column:resume_id; index" json:"resumeId,omitempty"`
	Type        string  `gorm:"column:type" json:"type"`
	Name        string  `gorm:"column:name" json:"name"`
	Description string  `gorm:"column:description" json:"description"`
	Percentage  float64 `gorm:"column:percentage" json:"percentage"`
}

// TableName for skill
func (s *Skill) TableName() string {
	return "resume_skills"
}
