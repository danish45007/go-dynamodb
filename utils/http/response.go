package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

func newResponse(data interface{}, status int) *response {
	return &response{
		Status: status,
		Result: data,
	}
}

func (res *response) bytes() []byte {
	data, _ := json.Marshal(res.Result)
	return data
}

func (res *response) string() string {
	return string(res.bytes())
}

func (res *response) sendResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(res.Status)
	w.Write(res.bytes())
	log.Println(res.string())
}

// status:200
func StatusOk(w http.ResponseWriter, r *http.Request, data interface{}) {
	newResponse(data, http.StatusOK).sendResponse(w, r)
}

// status:204
func StatusNoContent(w http.ResponseWriter, r *http.Request) {
	newResponse(nil, http.StatusNoContent).sendResponse(w, r)
}

// status:400
func StatusBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusBadRequest).sendResponse(w, r)
}

// status:404
func StatusNotFound(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusNotFound).sendResponse(w, r)
}

// status:405
func StatusMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	newResponse(nil, http.StatusMethodNotAllowed).sendResponse(w, r)
}

// status:409
func StatusConflict(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusConflict).sendResponse(w, r)
}

// status:500
func StatusInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusInternalServerError).sendResponse(w, r)
}
