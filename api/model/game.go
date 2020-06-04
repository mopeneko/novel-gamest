package model

// Game is an alternative object for Game entity
type Game struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	IsR18     bool   `json:"is_r18"`
	URL       string `json:"url"`
}
