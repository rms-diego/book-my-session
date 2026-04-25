package shared

type IDParam struct {
	ID string `uri:"id" validate:"required|uuid" message:"id is required and must be a valid UUID"`
}
