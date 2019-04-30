package domain

// Profile model
type Profile struct {
	ID        int    `gorm:"column:id; primary_key:yes" json:"id"`
	ResumeID  int    `gorm:"column:resume_id; index" json:"resumeId"`
	Fullname  string `gorm:"column:full_name" json:"fullName"`
	Religion  string `gorm:"column:religion" json:"religion"`
	Hobby     string `gorm:"column:hobby" json:"hobby"`
	Github    string `gorm:"column:github" json:"github"`
	Linkedin  string `gorm:"column:linkedin" json:"linkedin"`
	Instagram string `gorm:"column:instagram" json:"instagram"`
	Facebook  string `gorm:"column:facebook" json:"facebook"`
}

// TableName for profile
func (p *Profile) TableName() string {
	return "resume_profiles"
}
