package handlers

import (
	"errors"
	"net/http"

	"github.com/ViniciusdSC/mighty-blade-system-api/internal/api/models"
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/api/services"
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/presenters"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	ItemHandler interface {
		Route(r *gin.Engine)
	}

	itemHandler struct {
		itemService services.ItemService
	}
)

func NewItemHandler(itemService services.ItemService) ItemHandler {
	return &itemHandler{
		itemService,
	}
}

func (ih itemHandler) Route(r *gin.Engine) {
	r.GET("/item/:id", ih.addItemModelToContext, handlerToRoute(ih.getItem))
	r.GET("/item", handlerToRoute(ih.listItems))
	r.POST("/item", handlerToRoute(ih.createItem))
	r.PATCH("/item/:id", ih.addItemModelToContext, handlerToRoute(ih.updateItem))
	r.DELETE("/item/:id", ih.addItemModelToContext, handlerToRoute(ih.deleteItem))
}

func (ihd itemHandler) getItem(c *gin.Context) (*handlerResponse[models.ItemModel], error) {
	model, _ := c.Get("model")

	body := model.(*models.ItemModel)

	return &handlerResponse[models.ItemModel]{
		status: http.StatusOK,
		body:   body,
	}, nil
}

func (ihd itemHandler) listItems(c *gin.Context) (*handlerResponse[presenters.PresenterPagination[models.ItemModel]], error) {
	var filter services.ItemFilter
	err := c.BindQuery(&filter)

	if err != nil {
		return nil, err
	}

	if filter.Page == nil {
		filter.Page = new(int)
		*filter.Page = 1
	}

	if filter.PerPage == nil {
		filter.PerPage = new(int)
		*filter.PerPage = 10
	}

	count, err := ihd.itemService.Count(filter)
	if err != nil {
		return nil, err
	}

	items, err := ihd.itemService.Find(filter)
	if err != nil {
		return nil, err
	}

	return &handlerResponse[presenters.PresenterPagination[models.ItemModel]]{
		status: http.StatusOK,
		body: &presenters.PresenterPagination[models.ItemModel]{
			Page:    filter.Page,
			PerPage: filter.PerPage,
			Total:   *count,
			Items:   items,
		},
	}, nil
}

func (ihd itemHandler) createItem(c *gin.Context) (*handlerResponse[any], error) {
	var im models.ItemModel

	c.Bind(&im)

	err := ihd.itemService.Validate(&im)

	if err != nil {
		return nil, err
	}

	dbErr := ihd.itemService.Create(&im)

	if dbErr != nil {
		return nil, err
	}

	return &handlerResponse[any]{
		status: http.StatusCreated,
	}, nil
}

func (ihd itemHandler) updateItem(c *gin.Context) (*handlerResponse[any], error) {
	var im models.ItemModel

	c.Bind(&im)

	err := ihd.itemService.Validate(&im, "required")
	if err != nil {
		return nil, err
	}

	model, _ := c.Get("model")

	dbErr := ihd.itemService.Update(model.(*models.ItemModel), &im)
	if dbErr != nil {
		return nil, dbErr
	}

	return &handlerResponse[any]{
		status: http.StatusOK,
	}, nil
}

func (ihd itemHandler) deleteItem(c *gin.Context) (*handlerResponse[any], error) {
	model, _ := c.Get("model")

	err := ihd.itemService.Delete(model.(*models.ItemModel))
	if err != nil {
		return nil, err
	}

	return &handlerResponse[any]{
		status: http.StatusOK,
	}, nil
}

func (ihd itemHandler) addItemModelToContext(c *gin.Context) {
	id, exists := c.Params.Get("id")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "request should have ID in their path",
		})
		return
	}

	model, err := ihd.itemService.FindOne(id)

	if err == nil {
		c.Set("model", model)
		c.Next()
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusInternalServerError, err)
}
