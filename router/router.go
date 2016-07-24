package router

import (
	"gopkg.in/macaron.v1"

	"github.com/containerops/pilotage/handler"
)

//Pilotage Router Definition
func SetRouters(m *macaron.Macaron) {
	// Web API
	m.Get("/", handler.IndexV1Handler)
}
