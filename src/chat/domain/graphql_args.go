package domain

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared/filter"
)

// Param model
type Param struct {
	Keyword string
	GroupID int32
	filter.Filter
}
