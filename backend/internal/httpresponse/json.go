package httpresponse

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type envelope struct {
	Success bool       `json:"success"`
	Data    any        `json:"data,omitempty"`
	Error   *errorBody `json:"error,omitempty"`
}

type errorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Success(w http.ResponseWriter, status int, data any) {
	writeJSON(w, status, envelope{
		Success: true,
		Data:    data,
	})
}

func Error(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, envelope{
		Success: false,
		Error: &errorBody{
			Code:    code,
			Message: message,
		},
	})
}

func Decode(r *http.Request, dst any) error {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		return err
	}
	if decoder.More() {
		return errors.New("request body must contain a single JSON object")
	}
	return nil
}

func writeJSON(w http.ResponseWriter, status int, payload envelope) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, strings.TrimSpace(err.Error()), http.StatusInternalServerError)
	}
}
