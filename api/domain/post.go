package domain

// Post of the service
type Post struct {
	PostID string
	Text   string
	Author User
	Game   Game
}
