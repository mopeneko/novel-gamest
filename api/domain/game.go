package domain

// Game is infomation of novel games
type Game struct {
	GameID    string
	Title     string
	Thumbnail string
	IsR18     bool
	URL       string
}
