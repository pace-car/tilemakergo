package main

import "math"

func ColumnFromLongitude(lon float64, zoom int) int {
	return int(ColumnFromLongitudeF(lon, zoom))
}

func RowFromLatitude(lat float64, zoom int) int {
	return int(RowFromLatitudeF(lat, zoom))
}

func ColumnFromLongitudeF(lon float64, zoom int) float64 {
	return (lon + 180.0) / 360.0 * float64(math.Pow(2.0, float64(zoom)))
}

func RowFromLatitudeF(lat float64, zoom int) float64 {
	return float64(
		(1.0 - math.Log(math.Tan(float64(lat) * math.Pi / 180.0) + 1.0 / math.Cos(float64(lat) * math.Pi / 180.0)) / math.Pi) / 
		(2.0 * math.Pow(2.0, float64(zoom)))
	)
}
