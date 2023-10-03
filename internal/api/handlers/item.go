package handlers

import "github.com/gin-gonic/gin"

type (
	ItemHandler interface {
		getItem(c *gin.Context)
		listItems(c *gin.Context)
		createItem(c *gin.Context)
		updateItem(c *gin.Context)
		deleteItem(c *gin.Context)
	}

	itemHandler struct {
	}
)

func AddInRouter(r *gin.Engine) {
	ih := NewItemHandler()

	r.GET("/item", ih.listItems)
	r.GET("/item/:id", ih.getItem)
	r.POST("/item", ih.createItem)
	r.PUT("/item/:id", ih.updateItem)
	r.DELETE("/item/:id", ih.deleteItem)
}

func NewItemHandler() ItemHandler {
	return &itemHandler{}
}

func (ihd itemHandler) getItem(c *gin.Context) {
	c.JSON(200, gin.H{
		"item": true,
	})
}

func (ihd itemHandler) listItems(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": true,
	})
}

func (ihd itemHandler) createItem(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": true,
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
