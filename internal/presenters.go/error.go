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
