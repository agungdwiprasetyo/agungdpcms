package domain

// Resume main domain
type Resume struct {
	ID           int            `gorm:"column:id; primary_key:yes" json:"id"`
	Slug         string         `gorm:"column:slug; unique_index" json:"slug"`
	Name         string         `gorm:"column:name" json:"name"`
	Achievements []*Achievement `json:"achievements"`
	Educations   []*Education   `json:"educations"`
	Experiences  []*Experience  `json:"experiences"`
	Skills       []*Skill       `json:"skills"`
}

// TableName for resume domain
func (r *Resume) TableName() string {
	return "resumes"
}
