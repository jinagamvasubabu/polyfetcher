package schema

//status of OSM Request, this will take care of each and every area result and error if there are any
type OSMStatus struct {
	Error  error
	Result map[string]interface{}
}
