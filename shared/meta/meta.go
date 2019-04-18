package meta

type Meta struct {
	Page         int `json:"page"`
	Limit        int `json:"limit"`
	Offset       int `json:"-"`
	TotalPages   int `json:"totalPages"`
	TotalRecords int `json:"totalRecords"`
}

func (m *Meta) CalculateOffset() {
	m.Offset = (m.Page - 1) * m.Limit
}
