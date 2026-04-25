package filmsdto

type UpdateFilmRequest struct {
	Title       *string `json:"title" db:"title" validate:"min_len:1" message:"title must be at least 1 character long if provided"`
	Description *string `json:"description" db:"description"`
	Duration    *int    `json:"duration" db:"minutes_duration" validate:"gt:0" message:"duration must be greater than 0 if provided"`
	Language    *string `json:"language" db:"language" validate:"in:subtitled,dubbed" message:"language must be either 'subtitled' or 'dubbed' if provided"`
	ReleaseYear *int    `json:"releaseYear" db:"release_year" validate:"gt:0" message:"release year must be greater than 0 if provided"`
	Genre       *string `json:"genre" db:"genre" validate:"min_len:1" message:"genre must be at least 1 character long if provided"`
	AgeRange    *int    `json:"ageRange" db:"age_range"`
}
