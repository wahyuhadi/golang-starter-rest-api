/**
	Rahmat wahyu hadi
**/

package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response ...
type Response struct {
	ResponseWriter http.ResponseWriter
}

func (r *Response) SendOK(body interface{}) {
	setJSON(r.ResponseWriter)
	fmt.Println(body)
	encodeJSON(r.ResponseWriter, body)
}

func (r *Response) SendCreated(body interface{}) {
	setJSON(r.ResponseWriter)
	setHTTPStatus(r.ResponseWriter, http.StatusCreated)
	encodeJSON(r.ResponseWriter, body)
}

func (r *Response) SendNoContent() {
	setJSON(r.ResponseWriter)
	setHTTPStatus(r.ResponseWriter, http.StatusNoContent)
	encodeJSON(r.ResponseWriter, nil)
}

func (r *Response) SendBadRequest(message string) {
	http.Error(r.ResponseWriter, message, http.StatusBadRequest)
}

func (r *Response) SendNotFound() {
	http.Error(r.ResponseWriter, "Not Found", http.StatusNotFound)
}

func (r *Response) SendNotImplemented() {
	http.Error(r.ResponseWriter, "Not Implemented", http.StatusNotImplemented)
}

func setHTTPStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func setJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func encodeJSON(w http.ResponseWriter, body interface{}) {
	json.NewEncoder(w).Encode(body)
}
