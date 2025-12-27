package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/mr-raj2001/students-api/internal/types"
	"github.com/mr-raj2001/students-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

        var student types.Student  //creating an instance of student struct from types package

		err := json.NewDecoder(r.Body).Decode(&student)  //r.io Reader is passed to NewDecoder to read the request body  //Body type is io.ReadCloser which implements Reader interface
        
		// error has Is to check the type of errr 
		if errors.Is(err, io.EOF) {
			response.Writejson(w, http.StatusBadRequest, err.Error())
			return 
		}

		slog.Info("creating a student")
		//in go we have to serialize the data in request body to struct
		response.Writejson(w, http.StatusCreated, map[string]string{"success": "student created successfully"})  //sending response in json format with status code 201 using response package made by us
	}
}