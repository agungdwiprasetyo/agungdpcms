package domain

import "time"

// Education model
type Education struct {
	ID          int        `gorm:"column:id; primary_key:yes" json:"id"`
	ResumeID    int        `gorm:"column:resume_id; index" json:"resumeId"`
	Institution string     `gorm:"column:institution" json:"institution"`
	Degree      string     `gorm:"column:degree" json:"degree"`
	StartDate   *time.Time `gorm:"column:start_date" json:"startDate"`
	EndDate     *time.Time `gorm:"column:end_date" json:"endDate"`
}

// TableName for education
func (e *Education) TableName() string {
	return "resume_educations"
}
