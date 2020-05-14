package responses

import (
	"encoding/json"
	"net/http"
)

type responseData struct {
	Items  []interface{} `json:"items"`
	Count  int64         `json:"count"`
	Offset int64         `json:"offset"`
}

type successResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    responseData `json:"data"`
}

func writeJSON(res http.ResponseWriter, status int, data interface{}) error {
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

// SuccessJSON todo
func SuccessJSON(res http.ResponseWriter, status int, data interface{}) error {
	return writeJSON(res, status, &successResponse{
		Status:  "success",
		Message: "",
		Data: responseData{
			Items:  []interface{}{data},
			Count:  1,
			Offset: 0,
		},
	})
}
