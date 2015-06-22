package easyapi



import (
	"fmt"
	"time"
    "net/http"
    "encoding/json"
    "reflect"
	"io/ioutil"
)



type EasyContext struct {
	start, finish time.Time
	duration int
	method string
	reqBody []byte
}



type ResponseHolder struct {
	Status string `json:"status"`
	Response interface{} `json:"response"`
}



func(this *EasyContext) Process(w http.ResponseWriter, r *http.Request, controller ControllerInterface) {
	var (
		err *ApiError
		response ResponseHolder
		str string
		raw []byte
		e error
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
	if err != nil {
		if raw, e = json.Marshal(err); e == nil {
			str = string(raw)
		}
	} else {
		switch response.Response.(type) {
			case string:
				str = response.Response.(string)
			default:
				if raw, e = json.Marshal(response); e == nil {
					str = string(raw)
				}				
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch response.Response.(type) {
		case string:

		default:
			w.Header().Set("Content-Type", "application/json;charset=utf-8")
	}

	fmt.Fprintf(w, str)

	this.finishTime()
	this.log()
}



func(this *EasyContext) startTime() {
	this.start = time.Now()
}



func(this *EasyContext) finishTime() {
	this.finish = time.Now()
	this.duration = int(this.finish.Sub(this.start)/time.Millisecond)
}



func(this *EasyContext) readBody(r *http.Request, controller *ControllerInterface) (*ApiError) {
	var (
		e error
		)

	this.reqBody, e = ioutil.ReadAll(r.Body)
	if e != nil {
		return NewError("server error", "Can't read request", e.Error())
	}
 
 	if len(this.reqBody) > 0 {
	 	e = json.Unmarshal(this.reqBody, controller)
		if e != nil {
			return NewError("wrong format", "Request body format is wrong", e.Error())
		}
 	}	

 	return nil
}



func(this *EasyContext) log() {
	fmt.Println("")
	fmt.Println("[", this.start.Format("15:04:05"), "]")
	fmt.Println("METHOD:", this.method)
	//fmt.Println("REQUEST BODY", string(this.reqBody))
	fmt.Println("DURATION:", this.duration, "Milliseconds")
}