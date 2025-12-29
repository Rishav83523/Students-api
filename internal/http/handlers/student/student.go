package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mr-raj2001/students-api/internal/types"
	"github.com/mr-raj2001/students-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

        var student types.Student  //creating an instance of student struct from types package

		err := json.NewDecoder(r.Body).Decode(&student)  //r.io Reader is passed to NewDecoder to read the request body  //Body type is io.ReadCloser which implements Reader interface
        	slog.Info("creating a student")
		// error has Is to check the type of errr 
		if errors.Is(err, io.EOF) {
			response.Writejson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body"))) //if request body is empty we get EOF error so we send bad request response
			return 
		}

		if err != nil {
			response.Writejson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("invalid json body: %w", err))) //if there is any other error in decoding we send bad request response
			return
		}

		//validate request by using request validator package playground
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors) //type assertion to get validation errors
			response.Writejson(w, http.StatusBadRequest, response.ValidationError(validateErrs) )
			return
		}
	
		//in go we have to serialize the data in request body to struct
		response.Writejson(w, http.StatusCreated, map[string]string{"success": "student created successfully"})  //sending response in json format with status code 201 using response package made by us
	}
}