package render

import (
	"encoding/json"
	"net/http"
)

type M map[string]interface{}

type Binder interface {
	Bind(*http.Request) error
}

// PlainText writes a string to the response, setting the Content-Type as
// text/plain.
func PlainText(w http.ResponseWriter, status int, value string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte(value))
}

func JSON(w http.ResponseWriter, status int, value interface{}) {
	b, err := json.Marshal(value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(b)
}

func Bind(r *http.Request, v Binder) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	if err := v.Bind(r); err != nil {
		return err
	}
	return nil
}
