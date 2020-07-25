package polyfuse

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeometryUtils_CombinePolygons_Success(t *testing.T) {
	//Given
	gm := geometryUtils{logLevel: logrus.DebugLevel}
	areas := []string{"frankfurt", "munich"}

	//When
	response, err := gm.CombinePolygons(context.Background(), areas)

	//Then
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestGeometryUtils_CombinePolygons_areas_one_is_Invalid(t *testing.T) {
	//Given
	gm := geometryUtils{logLevel: logrus.DebugLevel}
	areas := []string{"asdasdasdasdasdasdas", "munich"}

	//When
	response, err := gm.CombinePolygons(context.Background(), areas)

	//Then
	assert.Nil(t, err)
	assert.Equal(t, Multipolygon, response.Type) //Because "Munich" Polygon type is Multipolygon
}

func TestGeometryUtils_CombinePolygons_failure(t *testing.T) {
	//Given
	gm := geometryUtils{logLevel: logrus.DebugLevel}
	areas := []string{"asdasdasdasdasdasdas", "asdasdasdasdasd"}

	//When
	_, err := gm.CombinePolygons(context.Background(), areas)

	//Then
	assert.NotNil(t, err)
}

func TestGeometryUtils_CombinePolygons_geometry_type_MultiPolygon(t *testing.T) {
	//Given
	gm := geometryUtils{logLevel: logrus.DebugLevel}
	areas := []string{"frankfurt", "munich"}

	//When
	response, err := gm.CombinePolygons(context.Background(), areas)

	//Then
	assert.Nil(t, err)
	assert.Equal(t, Multipolygon, response.Type)
}
