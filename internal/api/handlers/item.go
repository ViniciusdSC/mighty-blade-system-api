package handlers

import (
	"net/http"

	"github.com/ViniciusdSC/mighty-blade-system-api/internal/api/models"
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/api/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
			"message": "Something is wrong",
		})
		return
	}

	u, _ := uuid.Parse(id)

	model, err := ihd.itemService.FindOne(u)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "test",
		})
	}

	c.JSON(200, gin.H{
		"id": model,
	})
	return

	// u, err := uuid.Parse(id)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "Something is wrong",
	// 	})
	// 	return
	// }

	// model, err := ihd.itemService.FindOne(u)

	// c.JSON(200, model)
	// return
}

func (ihd itemHandler) listItems(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": true,
	})
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
	return
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
