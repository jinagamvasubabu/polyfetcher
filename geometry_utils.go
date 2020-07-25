package polyfuse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinagamvasubabu/polyfuse/schema"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

/**
// Program to fetch the polygon data of one or more areas
//  Not only that, using this you can combine two or more polygons/multipolygon to a Multi polygon
//
//  returns polygon if both areas are polygons
//  returns Multipolygon if both areas are multi polygons
//
//  validate the response using geojsonlint.com or geojson.io. Paste it directly in the textbox to validate it
//
//  Note: If you get an error like Polygons and MultiPolygons should follow the right-hand rule then follow this below article to fix it.
//
//  https://dev.to/jinagamvasubabu/solution-polygons-and-multipolygons-should-follow-the-right-hand-rule-2c8i
//
//  Errors:
//     * If one of the area is invalid area then it can fetch the other area if its valid
//     * If both are invalid then it will fail using
//
//  GoRoutines Support:
//     * Instead of getting each and every polygon data synchronously, Time metrics has been logged at the end of the program
//
*/

//IGeometryUtils Program to fetch the polygon data of one or more areas
type IGeometryUtils interface {
	CombinePolygons(ctx context.Context, areas []string) (string, error)
}

//GeometryUtils struct to configure GeometryUtils
type GeometryUtils struct {
	LogLevel log.Level
}

//CombinePolygons function to combine polygons or multipolygons
func (g *GeometryUtils) CombinePolygons(ctx context.Context, areas []string) (string, error) {
	log.SetLevel(g.LogLevel)
	defer calculateTimeTaken(time.Now(), "Time Taken by Fetch Polygons")
	logger := log.WithContext(ctx).WithFields(log.Fields{"Method": "CombinePolygons"})
	logger.Infof("Combine polygons for=%s", areas)
	logger.Debugf("GoRoutines count at beginning: %d", runtime.NumGoroutine())
	response := schema.GeoJson{}
	response.Coordinates = make([]interface{}, CoordinatesMaxLength) //Initialize with Max number of coordinates array
	respList, err := getPolygonDataFromOSM(ctx, areas)
	if err != nil {
		logger.WithFields(log.Fields{"err": err.Error()}).Error("error while fetching the polygon")
		return "", err
	}
	for i, result := range respList {
		geoJSON := result["geojson"].(map[string]interface{})
		coordinates := geoJSON["coordinates"].([]interface{})
		response.Type = geoJSON["type"].(string)
		//append based on the type of geojson
		if len(areas) == 1 {
			response.Coordinates = coordinates
		} else if i == 0 && response.Type == Polygon {
			response.Coordinates[0] = coordinates
		} else if i > 0 && response.Type == Polygon {
			count := lenOfCoOrdinatesArray(response.Coordinates)
			response.Coordinates[count] = coordinates
		} else if response.Type == Multipolygon {
			count := lenOfCoOrdinatesArray(response.Coordinates)
			for j := range coordinates {
				response.Coordinates[count] = coordinates[j]
				count++
			}
		}
	}
	if len(areas) > 1 {
		response.Type = Multipolygon
	}
	removeNilsFromArray(&response)
	logger.Debugf("GoRoutines count at last:%d", runtime.NumGoroutine())
	byteArr, err := json.Marshal(response)
	if err != nil {
		logger.WithFields(log.Fields{"err": err.Error()}).Errorf("Failed while marshalling the response", runtime.NumGoroutine())
		return "", err
	}
	return string(byteArr), nil
}

func getPolygonDataFromOSM(ctx context.Context, areas []string) ([]map[string]interface{}, error) {
	logger := log.WithContext(ctx).WithFields(log.Fields{"Method": "getPolygonDataFromOSM"})
	logger.Infof("getPolygonDataFromOSM for=%s", areas)
	var OSMData []map[string]interface{}
	c := make(chan schema.OSMStatus)
	noOfGoroutines := 0
	for _, v := range areas {
		go fetchOSMDataFromExternalClient(v, c)
		noOfGoroutines++
	}
	osmStatus := schema.OSMStatus{}
	isThereAnyErrorInGoRoutines := false
	for i := 0; i < noOfGoroutines; i++ {
		osmStatus = <-c
		if osmStatus.Error != nil {
			isThereAnyErrorInGoRoutines = true
		} else {
			OSMData = append(OSMData, osmStatus.Result)
		}
	}
	if isThereAnyErrorInGoRoutines && len(OSMData) == 0 {
		logger.Error("error while fetching the polygon ")
		return nil, errors.New("error while fetching the polygon")
	}
	return OSMData, nil
}

func fetchOSMDataFromExternalClient(area string, c chan schema.OSMStatus) {
	resp, err := http.Get(fmt.Sprintf(OsmURL, area))
	if err != nil {
		c <- schema.OSMStatus{Error: errors.New("error while fetching the polygon from OSM")}
	} else {
		defer resp.Body.Close()
	}
	if resp != nil && resp.Body != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c <- schema.OSMStatus{Error: errors.New("no data found in OSM")}
			return
		}
		var results []map[string]interface{}
		err = json.Unmarshal(body, &results)
		if err != nil {
			c <- schema.OSMStatus{Error: errors.New("no data found in OSM")}
			return
		}
		if len(results) == 0 {
			c <- schema.OSMStatus{Error: errors.New("no Data available in OSM")}
			return
		}
		c <- schema.OSMStatus{Result: results[0]} //always take the first result of it, because that's more relevant to the search
	}

}

func lenOfCoOrdinatesArray(coOrdinates []interface{}) int32 {
	var coOrdinatesLen int32
	for _, v := range coOrdinates {
		if v == nil {
			return coOrdinatesLen
		}
		coOrdinatesLen++
	}
	return coOrdinatesLen
}

func removeNilsFromArray(response *schema.GeoJson) {
	numOfCoordinates := 0
	for _, v := range response.Coordinates {
		if v == nil {
			break
		}
		numOfCoordinates++
	}
	coordinates := make([]interface{}, numOfCoordinates)
	coordinates = response.Coordinates[:numOfCoordinates]
	response.Coordinates = coordinates
}

func calculateTimeTaken(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
