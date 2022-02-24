package helper

import (
	"encoding/json"
	"io"
	"log"
)

func GetRequestBody(i io.ReadCloser) map[string]interface{} {
	payload := map[string]interface{}{}

	err := json.NewDecoder(i).Decode(&payload)
	if err != nil {
		log.Println(err)
	}

	return payload
}
