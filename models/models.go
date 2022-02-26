package models

type JsonResponse struct {
	Type, Data, Message interface{}
}

type GPSCoordinates struct {
	Lat, Lon float64
}
