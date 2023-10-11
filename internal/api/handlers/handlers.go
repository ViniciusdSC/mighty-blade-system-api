package handlers

import (
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/api/services"
)

type AppHandlers struct {
	IH ItemHandler
}

func CreateHandlers(itemService services.ItemService) AppHandlers {
	itemHandler := NewItemHandler(itemService)

	return AppHandlers{
		IH: itemHandler,
	}
}
