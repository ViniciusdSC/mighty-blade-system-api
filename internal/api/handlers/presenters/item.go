package presenters

type (
	CreateItemBody struct {
		Name string `json:"name" biding:"required"`
	}
)
