package polyfuse

const (
	CoordinatesMaxLength = 500
	Polygon              = "Polygon"
	MultiPolygon         = "MultiPolygon"
	OSM_URL              = "https://nominatim.openstreetmap.org/search.php?q=%s+&polygon_geojson=1&format=json" // nominatim is a UI for OSM data
)
