package filter

// Filter common model
type Filter struct {
	Page   int32  `json:"page"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
	Sort   string `json:"sort"`
	SortBy string `json:"sortBy"`
}

// CalculateOffset method
func (f *Filter) CalculateOffset() {
	f.Offset = (f.Page - 1) * f.Limit
}
