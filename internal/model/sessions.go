package model

type Session struct {
	ID        string  `json:"id" db:"id"`
	SeatLabel string  `json:"seat_label" db:"seat_label"`
	StartedAt string  `json:"started_at" db:"started_at"`
	EndedAt   string  `json:"ended_at" db:"ended_at"`
	UserID    string  `json:"user_id" db:"user_id"`
	FilmID    string  `json:"film_id" db:"film_id"`
	CreatedAt string  `json:"created_at" db:"created_at"`
	UpdatedAt string  `json:"updated_at" db:"updated_at"`
	DeletedAt *string `json:"deleted_at" db:"deleted_at"`
	Deleted   bool    `json:"deleted" db:"deleted"`
}

const SESSIONS_TABLE string = "sessions"
