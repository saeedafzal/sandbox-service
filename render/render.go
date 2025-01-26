package render

import "net/http"

func PlainText(w http.ResponseWriter, status int, value string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(value))
}
