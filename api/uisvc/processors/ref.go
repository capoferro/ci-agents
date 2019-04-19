package processors

import (
	"github.com/gin-gonic/gin"
	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/handlers"
)

func ListRefsByRepository(h *handlers.H, ctx *gin.Context) (interface{}, int, *errors.Error) {
	// user, err := getUser(h, ctx)
	// if err != nil {
	// 	return nil, 500, err
	// }
	refs, err := h.Clients.Data.GetRefsByRepository(ctx.GetString("repository"))
	if err != nil {
		return nil, 500, err
	}

	// TODO: ensure we don't expose refs from a repo that the user doesn't have access to.

	return refs, 200, nil
}
