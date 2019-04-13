package domain

import "time"

// Experience model
type Experience struct {
	ID        int        `gorm:"column:id; primary_key:yes" json:"id"`
	ResumeID  int        `gorm:"column:resume_id; index" json:"resumeId"`
	Title     string     `gorm:"column:title" json:"title"`
	Company   string     `gorm:"column:company" json:"company"`
	Location  string     `gorm:"column:location" json:"location"`
	StartDate *time.Time `gorm:"column:start_date" json:"startDate"`
	EndDate   *time.Time `gorm:"column:end_date" json:"endDate"`
}

// TableName for experience
func (e *Experience) TableName() string {
	return "resume_experiences"
}
