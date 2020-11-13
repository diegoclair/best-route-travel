package viewmodel

// TravelRequest viewmodel to handle a rest request
type TravelRequest struct {
	WhereFrom string `json:"where_from,omitempty"`
	WhereTo   string `json:"where_to,omitempty"`
}

// TravelResponse viewmodel to handle a rest response
type TravelResponse struct {
	Route string `json:"route,omitempty"`
	Price int64  `json:"price,omitempty"`
}
