package main

import (
	"math"
	"strconv"
)

// lat/lon to tile numbers
func tilenumber(lat, lon, zoom float64) (string, string) {
	n := math.Pow(2, zoom)
	xstr := strconv.FormatFloat((n*((lon+180)/360) - 1), 'f', 0, 64)
	n = math.Pow(2, zoom-1)
	ystr := strconv.FormatFloat((n * (1 - (math.Log(math.Tan(lat*(math.Pi/180))+1/(math.Cos(lat*(math.Pi/180)))) / math.Pi))), 'f', 0, 64)
	return xstr, ystr
}
