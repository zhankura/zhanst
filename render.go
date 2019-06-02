package zhanst

import (
	"encoding/json"
	"net/http"
)

type Render interface {
	Render(http.ResponseWriter, int) error
	WriteContentType(w http.ResponseWriter)
}

type JSON struct {
	Data interface{}
}

var jsonContentType = []string{"application/json; charset=utf-8"}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

func (r JSON) Render(w http.ResponseWriter, code int) (err error) {
	if err = WriteJSON(w, code, r.Data); err != nil {
		return err
	}
	return nil
}

func (r JSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

func WriteJSON(w http.ResponseWriter, code int, obj interface{}) error {
	writeContentType(w, jsonContentType)
	w.WriteHeader(code)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	if err != nil {
		return err
	}
	return nil
}
