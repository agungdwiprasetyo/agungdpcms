package domain

// Experience model
type Experience struct {
	ID        int    `gorm:"column:id; primary_key:yes" json:"id"`
	ResumeID  int    `gorm:"column:resume_id; index" json:"resumeId"`
	Title     string `gorm:"column:title" json:"title"`
	Company   string `gorm:"column:company" json:"company"`
	Location  string `gorm:"column:location" json:"location"`
	StartDate string `gorm:"column:start_date;type:timestamp with time zone" json:"startDate"`
	EndDate   string `gorm:"column:end_date;type:timestamp with time zone" json:"endDate"`
}

// TableName for experience
func (e *Experience) TableName() string {
	return "resume_experiences"
}
