/**
	Rahmat wahyu hadi
**/

package lib

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Request ...
type Request struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

// GetJSONBody ...
func (r *Request) GetJSONBody(model interface{}) {
	decoder := json.NewDecoder(r.Request.Body)
	decoder.Decode(&model)
}

// GetVarID ...
func (r *Request) GetVarID() (int, error) {
	vars := mux.Vars(r.Request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {

		http.Error(r.ResponseWriter, "Invalid id parameter", http.StatusBadRequest)
		return 0, err
	}
	return id, nil
}
