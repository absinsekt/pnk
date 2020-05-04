package responses

import (
	"encoding/json"
	"net/http"
)

// WriteJSON todo
func WriteJSON(res http.ResponseWriter, status int, data interface{}) error {
	raw, err := json.Marshal(data)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return err
	}

	res.WriteHeader(status)

	if _, err := res.Write(raw); err != nil {
		return err
	}

	return nil
}
