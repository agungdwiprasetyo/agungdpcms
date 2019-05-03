package domain

import (
	md "github.com/agungdwiprasetyo/agungdpcms/src/master/domain"
)

// Project model
type Project struct {
	ID          int            `gorm:"column:id; primary_key:yes" json:"id"`
	ResumeID    int            `gorm:"column:resume_id; index" json:"resumeId"`
	Name        string         `gorm:"column:name" json:"name"`
	Description string         `gorm:"column:description" json:"description"`
	Date        string         `gorm:"column:date;type:timestamp with time zone" json:"date"`
	URL         string         `gorm:"column:url" json:"url"`
	Languages   []*md.Language `gorm:"many2many:project_languages;" json:"languages"`
	Technology  string         `gorm:"column:technology" json:"technology"`
	Repository  string         `gorm:"column:repository" json:"repository"`
}

// TableName for project model
func (p *Project) TableName() string {
	return "resume_projects"
}
