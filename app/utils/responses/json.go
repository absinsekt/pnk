package responses

import (
	"encoding/json"
	"net/http"
)

// WriteJSON todo
func WriteJSON(w http.ResponseWriter, data interface{}) error {
	raw, err := json.Marshal(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	if _, err := w.Write(raw); err != nil {
		return err
	}

	return nil
}
