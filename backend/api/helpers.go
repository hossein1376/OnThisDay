package api

import (
	"encoding/json"
	"net/http"
	"time"
)

func writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Date", time.Now().Format(http.TimeFormat))

	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func writeStringJSON(w http.ResponseWriter, status int, data string, headers http.Header) error {
	js := []byte(data)
	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Date", time.Now().Format(http.TimeFormat))

	w.WriteHeader(status)
	w.Write(js)
	return nil
}
