package handlers

import (
	"net/http"

	"github.com/ViniciusdSC/mighty-blade-system-api/internal/api/models"
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/api/services"
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/presenters"
	"github.com/gin-gonic/gin"
)

type (
	ItemHandler interface {
		AddInRouter(r *gin.Engine)
		getItem(c *gin.Context)
		listItems(c *gin.Context)
		createItem(c *gin.Context)
		updateItem(c *gin.Context)
		deleteItem(c *gin.Context)
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

func (ih itemHandler) AddInRouter(r *gin.Engine) {
	r.GET("/item", ih.listItems)
	r.GET("/item/:id", ih.getItem)
	r.POST("/item", ih.createItem)
	r.PUT("/item/:id", ih.updateItem)
	r.DELETE("/item/:id", ih.deleteItem)
}

func (ihd itemHandler) getItem(c *gin.Context) {
	id, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "System was not able to get item ID",
		})
		return
	}

	model, err := ihd.itemService.FindOne(id)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "test",
		})
		return
	}

	c.JSON(200, gin.H{
		"id": model,
	})
}

func (ihd itemHandler) listItems(c *gin.Context) {
	var filter services.ItemFilter
	err := c.BindQuery(&filter)

	if err != nil {
		c.JSON(500, err)
		return
	}

	count, err := ihd.itemService.Count(filter)
	if err != nil {
		c.JSON(500, err)
		return
	}

	// items, err := ihd.itemService.Find(filter)
	// if err != nil {
	// 	c.JSON(500, err)
	// 	return
	// }

	response := presenters.PresenterPagination[models.ItemModel]{
		Page:    *filter.Page,
		PerPage: *filter.PerPage,
		Total:   *count,
		Items:   []models.ItemModel{},
	}

	c.JSON(200, response)
}

func (ihd itemHandler) createItem(c *gin.Context) {
	var im models.ItemModel

	c.Bind(&im)

	err := ihd.itemService.Validate(&im)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	dbErr := ihd.itemService.Create(&im)

	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, dbErr)
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"result": im,
	})
}

func (ihd itemHandler) updateItem(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": true,
	})
}

func (ihd itemHandler) deleteItem(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": true,
	})
}
