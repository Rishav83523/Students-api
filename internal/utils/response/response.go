package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Error string `json:"error"`  //basically we are saying that in json response the field name is error but in struct it is Errorit is seralizing using package encoding json
}

const (
	StausOk = "ok"
	StatusError = "error"
)

//to send response in json format
//data `any` it mans we dont know the type of data it can be struct , map , slice etc which we are going to write
func Writejson(w http.ResponseWriter, statusCode int, data any) error {
    w.Header().Set("Content-Type", "application/json") //setting the content type header to application json
	w.WriteHeader(statusCode) 
	
	return json.NewEncoder(w).Encode(data)//encoding the data to json and writing it to response writer it return error type 
}


func GeneralError( err error) Response {
     return Response{ 
		Status: StatusError,
		Error: err.Error(),
	 }
}