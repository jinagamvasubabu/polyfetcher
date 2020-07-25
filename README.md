# polyfuse : A Go Package to  combine two or more polygons into a single Polygon/MultiPolygon
![alt text](https://github.com/jinagamvasubabu/polyfuse/blob/master/images/polyfuse_brand.png?raw=true)
## Overview 
[![Build Status](https://travis-ci.org/jinagamvasubabu/polyfuse.svg?branch=master)](https://travis-ci.org/jinagamvasubabu/polyfuse)
[![Build Status](https://circleci.com/gh/jinagamvasubabu/polyfuse.svg?style=svg)](https://circleci.com/gh/jinagamvasubabu/polyfuse)
[![Go Report Card](https://goreportcard.com/badge/github.com/jinagamvasubabu/polyfuse)](https://goreportcard.com/report/github.com/jinagamvasubabu/polyfuse)
[![GoDoc](https://godoc.org/github.com/jinagamvasubabu/polyfuse?status.svg)](https://godoc.org/github.com/jinagamvasubabu/polyfuse) 
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Program to fetch the polygon data of one or more areas
Not only that, using this you can combine two or more polygons/multipolygon to a Multi polygon

returns polygon if both areas are polygons
  returns Multipolygon if both areas are multi polygons

  validate the response using geojsonlint.com or geojson.io. Paste it directly in the textbox to validate it

 Note: If you get an error like Polygons and MultiPolygons should follow the right-hand rule then follow this below article to fix it.

  https://dev.to/jinagamvasubabu/solution-polygons-and-multipolygons-should-follow-the-right-hand-rule-2c8i

 Errors:
   * If one of the area is invalid area then it can fetch the other area if its valid
     * If both are invalid then it will fail using

 GoRoutines Support:
     * Instead of getting each and every polygon data synchronously, Time metrics has been logged at the end of the program


## Install

```
go get github.com/jinagamvasubabu/polyfuse
```

## Author

JinagamVasubabu

## License

Apache 2.0.
