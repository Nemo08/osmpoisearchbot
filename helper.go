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

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// Distance function returns the distance (in meters) between two points of
//     a given longitude and latitude relatively accurately (using a spherical
//     approximation of the Earth) through the Haversin Distance Formula for
//     great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
//
// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func Round(f float64) int64 {
	return int64(math.Floor(f + .5))
}
