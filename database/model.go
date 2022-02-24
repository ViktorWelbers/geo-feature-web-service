package database

type Movie struct {
	MovieID   string `json:"movieid"`
	MovieName string `json:"moviename"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    string `json:"data"`
	Message string `json:"message"`
}
