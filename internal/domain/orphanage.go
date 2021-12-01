package domain

type Orphanage struct {
	Base
	Name           string  `json:"name"`
	About          string  `json:"about"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Instructions   string  `json:"instructions"`
	OpenedHours    string  `json:"opened_hours"`
	OpenOnWeekends bool    `json:"open_on_weekends"`
	Images         []Image `json:"images,omitempty"`
}
