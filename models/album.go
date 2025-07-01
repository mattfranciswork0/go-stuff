package models

type Album struct {
	ID     int32   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
