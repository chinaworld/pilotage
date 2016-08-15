package handler

import (
	"encoding/json"
	"net/http"

	"github.com/containerops/pilotage/models"

	"gopkg.in/macaron.v1"
)

//PostServiceDefinitionForm is
type PostServiceDefinitionForm struct {
	Service        string `from:"service" binding:"Required"`
	Title          string `from:"title" binding:"Required"`
	Gravatar       string `from:"gravatar" binding:""`
	Endpoints      string `from:"endpoints" binding:""`
	Environments   string `from:"environments" binding:""`
	Authorizations string `from:"authorizations" binding:""`
	Configurations string `from:"configurations" binding:""`
	Description    string `from:"description" binding:""`
}

//PostServiceDefinitionV1Handler is function intergrated the third DevOps service with Pilotage.
//Support:
//	- [Github](https://github.com) [TODO]
//	- [Slack](https://slack.com) [TODO]
//	- [BitBucket](https://bitbucket.org) [TODO]
//  - [Gitter](https://gitter.im) [TODO]
//	- [Travis CI](https://travis-ci.org/) [TODO]
func PostServiceDefinitionV1Handler(ctx *macaron.Context, s PostServiceDefinitionForm) (int, []byte) {
	df := models.ServiceDefinition{}

	if id, err := df.Create(s.Service, s.Title, s.Gravatar, s.Endpoints, s.Environments, s.Authorizations, s.Configurations, s.Description); err != nil {
		result, _ := json.Marshal(map[string]string{"message": err.Error()})
		return http.StatusBadRequest, result
	} else {
		result, _ := json.Marshal(map[string]interface{}{"service": s.Service, "id": id})
		return http.StatusOK, result
	}
}

//GetServiceDefinitionListV1Handler is
func GetServiceDefinitionListV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//PutServiceDefinitionV1Handler is
func PutServiceDefinitionV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//GetServiceDefinitionV1Handler is
func GetServiceDefinitionV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//DeleteServiceDefinitionV1Handler is
func DeleteServiceDefinitionV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//PostServiceV1Handler is
func PostServiceV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//PutServiceV1Handler is
func PutServiceV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//GetServiceV1Handler is
func GetServiceV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//DeleteServiceV1Handler is
func DeleteServiceV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}

//AnyServiceCallbackV1Handler is
func AnyServiceCallbackV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": ""})
	return http.StatusOK, result
}
