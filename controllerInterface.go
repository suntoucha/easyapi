package easyapi



import (
    "net/http"
)



type ControllerInterface interface {
	SetRequest(*http.Request)
	SetResponse(*http.ResponseWriter)
	Validate() (*ApiError)
	Payload() (interface{}, *ApiError)
}