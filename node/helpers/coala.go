package helpers

import (
	"encoding/json"
	"github.com/coalalib/coalago"
)

func OutputJson(code coalago.CoapCode, result interface{}) *coalago.CoAPResourceHandlerResult {
	jsonResult, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return OutputMessage(coalago.CoapCodeInternalServerError, "json parsing error")
	}
	return coalago.NewResponse(coalago.NewBytesPayload(jsonResult), code)
}

func OutputMessage(code coalago.CoapCode, message string) *coalago.CoAPResourceHandlerResult {
	return coalago.NewResponse(coalago.NewStringPayload(message), code)
}

func OutputData(code coalago.CoapCode, data []byte) *coalago.CoAPResourceHandlerResult {
	return coalago.NewResponse(coalago.NewBytesPayload(data), code)
}
