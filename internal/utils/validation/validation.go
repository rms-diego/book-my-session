package validation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"github.com/rms-diego/book-my-session/pkg/exception"
)

// BindAndValidateBody deserializes the JSON request body into req and runs the
// struct validation rules defined via `validate:` tags (gookit/validate).
// Returns the first error encountered: a binding error if the JSON is malformed
// or missing required fields, or a validation error if a rule is violated.
func BindAndValidateBody[T any](c *gin.Context, payloadStruct *T) error {
	if err := c.ShouldBindJSON(payloadStruct); err != nil {
		return exception.NewException(err.Error(), http.StatusBadRequest)
	}

	v := validate.Struct(payloadStruct)
	if !v.Validate() {
		return exception.NewException(v.Errors.One(), http.StatusBadRequest)
	}

	return nil
}

// BindAndValidateParams deserializes URI parameters into req and runs the
// struct validation rules defined via `validate:` tags (gookit/validate).
// Returns the first error encountered: a binding error if parameters are missing
// or of the wrong type, or a validation error if a rule is violated.
func BindAndValidateParams[T any](c *gin.Context, payloadStruct *T) error {
	if err := c.ShouldBindUri(payloadStruct); err != nil {
		return exception.NewException(err.Error(), http.StatusBadRequest)
	}

	v := validate.Struct(payloadStruct)
	if !v.Validate() {
		return exception.NewException(v.Errors.One(), http.StatusBadRequest)
	}

	return nil
}
