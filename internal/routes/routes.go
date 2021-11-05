package routes

import (
	"github.com/go-chi/chi"
)

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig().SetTimeOut(ServiceConfig.GetConfig().timeout),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouter() *chi.Mux {

}
func (r *Router) setConfigRouters() {

}

func RouterHealth() {

}

func ProductRouter() {

}

func EnableCORS() {}

func EnableRecover() {}

func EnableRequestID() {}

func EnableRealIP() {}
