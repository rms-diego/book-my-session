package filmsdto

type CreateFilmRequest struct {
	Title       string `json:"title" db:"title" validate:"required|min_len:1" message:"title is required"`
	Description string `json:"description" db:"description"`
	Duration    int    `json:"duration" db:"minutes_duration" validate:"required|gt:0" message:"duration must be greater than 0"`
	Language    string `json:"language" db:"language" validate:"required|in:subtitled,dubbed" message:"language must be either 'subtitled' or 'dubbed'"`
	ReleaseYear int    `json:"releaseYear" db:"release_year" validate:"required|gt:0" message:"release year is required or release year must be greater than 0"`
	Genre       string `json:"genre" db:"genre" validate:"required|min_len:1" message:"genre is required"`
	AgeRange    int    `json:"ageRange" db:"age_range"`
}
