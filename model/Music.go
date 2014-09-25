package model

// Music struct
type Music struct {
	ID       string `json:"id" db:"id"`
	ArtistID string `json:"artist_id" db:"artist_id"`
	Title    string `json:"title" db:"title"`
	Outline  string `json:"outline" db:"outline"`
}
