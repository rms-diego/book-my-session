package validation

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

// BindAndValidate deserializes the JSON request body into req and runs the
// struct validation rules defined via `validate:` tags (gookit/validate).
// Returns the first error encountered: a binding error if the JSON is malformed
// or missing required fields, or a validation error if a rule is violated.
func BindAndValidate[T any](c *gin.Context, req *T) error {
	if err := c.ShouldBindJSON(req); err != nil {
		return err
	}

	v := validate.Struct(req)
	if !v.Validate() {
		return errors.New(v.Errors.One())
	}

	return nil
}
