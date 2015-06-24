package easyapi

import (
	"net/http"
	"testing"
)

const (
	host = "google.com"
)

func TestSetRequest(t *testing.T) {
	ec := new(EasyController)
	req := new(http.Request)
	req.Host = host
	ec.SetRequest(req)
	if ec.Request.Host != host {
		t.Error("Ошибка, хосты не совпадают. Реквест либо не был установлен либо уже изменен")
	}
}

func TestSetResponse(t *testing.T) {
	ec := new(EasyController)
	res := new(http.ResponseWriter)
	ec.SetResponse(res)
	if ec.Response != res {
		t.Error("Ошибка, вектора на респонс-райтеры не совпадают. Райтер либо не был установлен, либо уже изменен")
	}
}
