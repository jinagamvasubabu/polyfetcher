package schema

// Open street map response structure
type GeoJson struct {
	Type        string        `json:"type"`
	Coordinates []interface{} `json:"coordinates"`
}
