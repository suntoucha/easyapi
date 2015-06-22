package easyapi



import (
    "net/http"
)



type EasyController struct {
	Request *http.Request
	Response *http.ResponseWriter
}



func(this *EasyController) SetRequest(r *http.Request) {
	this.Request = r
}



func(this *EasyController) SetResponse(w *http.ResponseWriter) {
	this.Response = w
}