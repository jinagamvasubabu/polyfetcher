package polyfuse

type GeoJson struct {
	Type string `json:"type"`
	Coordinates []interface{} `json:"coordinates"`
}