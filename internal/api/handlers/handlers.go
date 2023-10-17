package handlers

import (
	"errors"
	"net/http"

	"github.com/ViniciusdSC/mighty-blade-system-api/internal/api/services"
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/presenters"
	"github.com/gin-gonic/gin"
)

type AppHandlers struct {
	IH ItemHandler
}

type handlerResponse[T any] struct {
	status int
	body   *T
}

func NewAppHandler(s services.AppServices) AppHandlers {
	itemHandler := NewItemHandler(s.IS)

	return AppHandlers{
		IH: itemHandler,
	}
}

func handlerToRoute[T any](handler func(ctx *gin.Context) (*handlerResponse[T], error)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response, err := handler(ctx)

		if err == nil {
			if response.body == nil {
				ctx.Status(response.status)
				return
			}
			ctx.JSON(response.status, response.body)
			return
		}

		if errors.Is(err, presenters.NotFoundError{}) {
			ctx.Status(http.StatusNotFound)
			return
		}

		if errors.Is(err, presenters.ValidationError{}) {
			ve := err.(presenters.ValidationError)
			ctx.JSON(ve.Status, err)
			return
		}

		ctx.JSON(http.StatusInternalServerError, err)
	}
}
