package easyapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

type EasyContext struct {
	start, finish time.Time
	duration      int
	method        string
	ReqBody       []byte
}

type ResponseHolder struct {
	Status   string      `json:"status"`
	Response interface{} `json:"response"`
}

// completeResponse says to context to pass response from controller as is
type completeResponse struct{}

// CompleteResponse says to context to pass response from controller as is
var CompleteResponse = completeResponse{}

func (this *EasyContext) Process(w http.ResponseWriter, r *http.Request, controller ControllerInterface) {
	var (
		err      *ApiError
		response ResponseHolder
		str      string
		raw      []byte
		e        error
	)
	this.startTime()
	this.method = reflect.TypeOf(controller).String()
	controller.SetRequest(r)
	controller.SetResponse(&w)
	if err = this.readBody(r, &controller); err == nil {
		if err = controller.Validate(); err == nil {
			response.Status = "ok"
			response.Response, err = controller.Payload()
		}
	}
	str = ""
	switch response.Response.(type) {
	case completeResponse:
	case string:
		str = response.Response.(string)
	default:
		if err != nil {
			if raw, e = json.Marshal(err); e == nil {
				str = string(raw)
			}
		} else {
			if raw, e = json.Marshal(response); e == nil {
				//Marshal по умолчанию переводит символы <, > и & в коды
				//А нам это не нужно — поэтому возвращаем их обратно
				raw = bytes.Replace(raw, []byte("\\u003c"), []byte("<"), -1)
				raw = bytes.Replace(raw, []byte("\\u003e"), []byte(">"), -1)
				raw = bytes.Replace(raw, []byte("\\u0026"), []byte("&"), -1)
				str = string(raw)
			}
		}
	}
	switch response.Response.(type) {
	case completeResponse:
		this.finishTime()
		this.log()
		return
	case string:
	default:
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
	}
	fmt.Fprint(w, str)
	this.finishTime()
	this.log()
}

func (this *EasyContext) startTime() {
	this.start = time.Now()
}

func (this *EasyContext) finishTime() {
	this.finish = time.Now()
	this.duration = int(this.finish.Sub(this.start) / time.Millisecond)
}

func (this *EasyContext) readBody(r *http.Request, controller *ControllerInterface) *ApiError {
	var e error
	this.ReqBody, e = ioutil.ReadAll(r.Body)
	if e != nil {
		return NewError("server error", "Can't read request", e.Error())
	}
	if len(this.ReqBody) > 0 {
		e = json.Unmarshal(this.ReqBody, controller)
		if e != nil {
			return NewError("wrong format", "Request body format is wrong", e.Error())
		}
	}
	return nil
}

func (this *EasyContext) log() {
	fmt.Println("")
	fmt.Println("[", this.start.Format("15:04:05"), "]")
	fmt.Println("METHOD:", this.method)
	fmt.Println("DURATION:", this.duration, "Milliseconds")
}
