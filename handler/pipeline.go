package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/macaron.v1"
)

//PostPipelineV1Handler is create pipeline data with namespace/repository and basic pipeline data.
func PostPipelineV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//PostPipelineJSONV1Handler is create pipeline with entire data.
func PostPipelineJSONV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//GetPipelineV1Handler is get pipeline data, json/yaml format.
func GetPipelineV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//PutPipelineV1Handler is update the pipeline data, json format data in the http body.
func PutPipelineV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//DeletePipelineV1Handler is delete the pipeline data.
func DeletePipelineV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//ExecutePipelineV1Handler is
func ExecutePipelineV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//GetOutcomeListV1Handler is
func GetOutcomeListV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//GetOutcomeV1Handler is
func GetOutcomeV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}
