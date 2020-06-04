package model

// Post is an alternative object for Post entity
type Post struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Author User   `json:"author"`
	Game   Game   `json:"game"`
}
