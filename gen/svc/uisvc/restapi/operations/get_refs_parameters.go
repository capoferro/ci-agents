package operations

import (
	"github.com/gin-gonic/gin"
	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/handlers"
)

// GetRefsValidateURLParams validates the parameters in the
// URL according to the swagger specification.
func GetRefsValidateURLParams(h *handlers.H, ctx *gin.Context) *errors.Error {
	repository := ctx.Query("repository")

	if len(repository) == 0 {
		return errors.New("'/refs': parameter 'repository' is empty")
	}

	ctx.Set("repository", repository)

	return nil
}
