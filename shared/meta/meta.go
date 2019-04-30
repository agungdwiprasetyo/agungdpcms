package meta

import "math"

// Meta model
type Meta struct {
	Page         int `json:"page"`
	Limit        int `json:"limit"`
	TotalPages   int `json:"totalPages"`
	TotalRecords int `json:"totalRecords"`
}

// CalculatePages meta method
func (m *Meta) CalculatePages() {
	m.TotalPages = int(math.Ceil(float64(m.TotalRecords) / float64(m.Limit)))
}
