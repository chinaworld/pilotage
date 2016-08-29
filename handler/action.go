package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/macaron.v1"
)

//PostActionV1Handler is
func PostActionV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//GetActionV1Handler is
func GetActionV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//PutActionV1Handler is
func PutActionV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//DeleteActionV1Handler is
func DeleteActionV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//PutStartActionV1Handler is
func PutStartActionV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//PutExecuteActionV1Handler is
func PutExecuteActionV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//PutStatusActionV1Handler is
func PutStatusActionV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//PutResultActionV1Handler is
func PutResultActionV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//PutDeleteActionV1Handle is
func PutDeleteActionV1Handle(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}
