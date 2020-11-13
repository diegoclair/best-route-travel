package entity

// BestRoute entity
type BestRoute struct {
	Route string `json:"route,omitempty"`
	Price int64  `json:"price,omitempty"`
}

// Filedata entity
type Filedata struct {
	WhereFrom string `json:"where_from,omitempty"`
	WhereTo   string `json:"where_to,omitempty"`
	Price     int64  `json:"price,omitempty"`
}
