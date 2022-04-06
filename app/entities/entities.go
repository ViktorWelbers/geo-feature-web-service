package entities

type GPSCoordinates struct {
	Lat, Lon, Radius float64
}

type PassthroughObj struct {
	FeatureVector map[string]int `json:"data"`
}
