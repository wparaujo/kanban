package board

import (
	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/middleware"
	"net/http"
)

// ListBoards gets a list of board accessible by the authenticated user.
func ListBoards(ctx *middleware.Context) {
	boards, err := ctx.DataSource.ListBoards()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &models.Response{
		Data: boards,
	})
}

// ItemBoard gets a specific board, identified by project ID or
// NAMESPACE/BOARD_NAME, which is owned by the authenticated user.
func ItemBoard(ctx *middleware.Context) {
	board, err := ctx.DataSource.ItemBoard(ctx.Query("project_id"))

	if err != nil {
		if err, ok := err.(models.ReceivedDataErr); ok {
			ctx.JSON(err.StatusCode, &models.ResponseError{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &models.Response{
		Data: board,
	})
}

func Configure(ctx *middleware.Context, form models.BoardRequest) {
	status, err := ctx.DataSource.ConfigureBoard(&form)

	if err != nil {
		ctx.JSON(status, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &models.Response{})
}
