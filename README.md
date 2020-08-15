polyfetcher
--
A Go Package can fetch polygon definitions from OpenStreet maps and it can also combine one or more polygons into one.
![alt text](https://github.com/jinagamvasubabu/polyfetcher/blob/master/images/poly_fetcher.jpg?raw=true)
## Overview 
[![Build Status](https://travis-ci.org/jinagamvasubabu/polyfetcher.svg?branch=master)](https://travis-ci.org/jinagamvasubabu/polyfetcher)
[![Build Status](https://circleci.com/gh/jinagamvasubabu/polyfetcher.svg?style=svg)](https://circleci.com/gh/jinagamvasubabu/polyfetcher)
[![Go Report Card](https://goreportcard.com/badge/github.com/jinagamvasubabu/polyfetcher)](https://goreportcard.com/report/github.com/jinagamvasubabu/polyfetcher)
[![GoDoc](https://godoc.org/github.com/jinagamvasubabu/polyfetcher?status.svg)](https://godoc.org/github.com/jinagamvasubabu/polyfetcher) 
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![HitCount](http://hits.dwyl.com/jinagamvasubabu/githubcom/jinagamvasubabu/polyfetcher.svg)](http://hits.dwyl.com/jinagamvasubabu/githubcom/jinagamvasubabu/polyfetcher)
[![Maintainability](https://api.codeclimate.com/v1/badges/bc775ba7e9b6231175a2/maintainability)](https://codeclimate.com/github/jinagamvasubabu/polyfetcher/maintainability)

## Documentation ?
More you can find here: [https://hangoutdude.com/polyfetcher](https://hangoutdude.com/polyfetcher/)

## What it does?
* It can fetch Polygon definitions from `OpenstreetMaps` by taking an area name, throw error if the area is not in OSM database
* Combine one or more polygons into one single `Multipolygon`

## Install

```
go get github.com/jinagamvasubabu/polyfetcher
```



## How to use Fetch Polygons?
* get `polyfetcher`
```
  go get github.com/jinagamvasubabu/polyfetcher
```
* import and use it like below:
```
  import "github.com/jinagamvasubabu/polyfetcher"
  import "github.com/sirupsen/logrus"
 
  p := polyfetcher.GeometryUtils{}
  resp, err := p.FetchPolygons(context.Background(), []string{"germany", "belgium"})
  fmt.Println(resp)
```

## How to use Combine Polygons?
* get `polyfetcher`
```
  go get github.com/jinagamvasubabu/polyfetcher
```
* import and use it like below:
```
  import "github.com/jinagamvasubabu/polyfetcher"
  import "github.com/sirupsen/logrus"
 
  p := polyfetcher.GeometryUtils{}
  resp, err := p.CombinePolygons(context.Background(), []string{"germany", "belgium"})
  fmt.Println(resp)
```


### Response:
```go
//GeoJson Open street map response structure
type GeoJson struct {
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	Coordinates []interface{} `json:"coordinates"`
}
```


## How to enable Debug logs?

* Pass logrus `LogLevel:logrus.DebugLevel` to `GeometryUtils` struct
```
  import "github.com/jinagamvasubabu/polyfetcher"

 
  p := polyfetcher.GeometryUtils{LogLevel:logrus.DebugLevel}
  resp, err := p.CombinePolygons(context.Background(), []string{"germany", "belgium"})
```

## How to validate the data ?
Response type is schema.GeoJson, you can marshal the response to string and use [geojsonlint](https://geojsonlint.com) to validate the geojson.
![alt text](https://github.com/jinagamvasubabu/polyfetcher/blob/master/images/geojsonlint.png?raw=true)

Note: 
If you get an error like Polygons and MultiPolygons should follow the right-hand rule then follow this below article to fix it.
[https://dev.to/jinagamvasubabu/solution-polygons-and-multipolygons-should-follow-the-right-hand-rule-2c8i](https://dev.to/jinagamvasubabu/solution-polygons-and-multipolygons-should-follow-the-right-hand-rule-2c8i)


## Known Errors: 
Errors:
   * If one of the area is invalid area then it can fetch the other area if its valid
   * If both are invalid then it will throw error

## Goroutines Support:
Goroutines support to fetch areas concurrently instead of synchronously

## Author

JinagamVasubabu

## License

Apache 2.0.
