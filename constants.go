package polyfuse

const (
	//CoordinatesMaxLength Maximum length of the coordinates array
	CoordinatesMaxLength = 500
	//Polygon constant Polygon
	Polygon = "Polygon"
	//Multipolygon constant MultiPolygon
	Multipolygon = "Multipolygon"
	//OsmURL Url to fetch the OSM data for an area
	OsmURL = "https://nominatim.openstreetmap.org/search.php?q=%s+&polygon_geojson=1&format=json" // nominatim is a UI for OSM data
)
