package presenters

type PresenterValidationErr struct {
	Tag string `json:"tag"`
	Key string `json:"key"`
}

type PresenterValidationErrors struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Errors  []PresenterValidationErr `json:"errors"`
}

type ValidationError struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Errors  []PresenterValidationErr `json:"errors"`
}

func (ve ValidationError) Error() string {
	return "Validation Error!"
}

type NotFoundError struct{}

func (nfe NotFoundError) Error() string {
	return "Entity not found!"
}
