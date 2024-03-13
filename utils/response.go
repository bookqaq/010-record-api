package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// wrapper for json response, body must be convertable to json when called
func ResponseJSON(w http.ResponseWriter, code int, body any) {
	w.WriteHeader(code)

	data, err := json.Marshal(body)
	if err != nil {
		panic(fmt.Errorf("response json marshal failed: %w", err))
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}
