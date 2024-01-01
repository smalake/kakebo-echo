package structs

type HttpResponse struct {
	Code    int
	Message string
	Error   error
	Data    interface{}
}
