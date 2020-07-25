package polyfuse

import "context"

type GeometryUtils interface {
	CombinePolygons(ctx context.Context, areas []string) (GeoJson, error)
}

type geometryUtils struct {

}

