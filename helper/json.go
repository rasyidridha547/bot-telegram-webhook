package helper

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func Decode(resp *resty.Response) map[string]interface{} {
	var response map[string]interface{}

	err := json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil
	}

	if response == nil {
		fmt.Println("error json unmarshal")
	}

	return response
}
