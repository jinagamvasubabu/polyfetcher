package schema

//GeoJson Open street map response structure
type GeoJson struct {
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	Coordinates []interface{} `json:"coordinates"`
}


