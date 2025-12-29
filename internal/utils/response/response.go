package response

import (
	"encoding/json"
	"fmt" // Importing the fmt package
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10" // Importing the validator package
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


func ValidationError( errs validator.ValidationErrors) Response { 
      var errMsgs []string

	  for _, e := range errs {
		switch e.ActualTag() {
		case "required":
			errMsgs = append(errMsgs,  fmt.Sprintf("The field %s is required", e.Field()))
		case "email":
			errMsgs = append(errMsgs,  fmt.Sprintf("The field %s must be a valid email address", e.Field()))    //e.Field() gives the field name
		}
	  }
     
	  return Response{
		Status: StatusError,
		Error: strings.Join(errMsgs, ", "),  //joining all the error messages with comma separator
	  }
}