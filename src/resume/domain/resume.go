package domain

type Resume struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Percentage float64 `json:"percentage"`
}
