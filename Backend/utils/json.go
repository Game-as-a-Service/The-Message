package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func RequestToJsonBody(requestBody any) []byte {
	bytesRepresentation, _ := json.Marshal(requestBody)
	return bytesRepresentation
}

func JsonBodyToMap(response *http.Response) map[string]interface{} {
	responseBodyAsByteArray, _ := io.ReadAll(response.Body)
	responseBody := make(map[string]interface{})
	_ = json.Unmarshal(responseBodyAsByteArray, &responseBody)
	return responseBody
}
