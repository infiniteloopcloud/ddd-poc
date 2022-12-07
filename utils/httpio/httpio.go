package httpio

import (
	"encoding/json"
	"io"
	"net/http"
)

func Bind[T any](body io.ReadCloser) (T, error) {
	defer body.Close()
	var t T
	if err := json.NewDecoder(body).Decode(&t); err != nil {
		return t, err
	}
	return t, nil
}

func ResponseSuccess(w http.ResponseWriter, v any) {
	respond(w, v, http.StatusOK)
}

func ResponseBadRequest(w http.ResponseWriter) {
	respond(w, nil, http.StatusBadRequest)
}

func respond(w http.ResponseWriter, v any, status int) {
	if v != nil {
		data, err := json.Marshal(v)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(data); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}
