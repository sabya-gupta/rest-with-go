package errors

type RestError struct {
	Message string
	Status  int
	Error   string
}
