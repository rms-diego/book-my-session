package model

type Film struct {
	ID              string `json:"id" db:"id"`
	Title           string `json:"title" db:"title"`
	Description     string `json:"description" db:"description"`
	MinutesDuration int    `json:"minutes_duration" db:"minutes_duration"`
	Language        string `json:"language" db:"language"`
	Thumbnail       string `json:"thumbnail" db:"thumbnail"`
	ReleaseDate     string `json:"release_date" db:"release_date"`
	CreatedAt       string `json:"created_at" db:"created_at"`
	UpdatedAt       string `json:"updated_at" db:"updated_at"`
}

const FILMS_TABLE string = "films"
