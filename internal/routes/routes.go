package routes

import (
	ServiceConfig "github.com/danish45007/go-dynamodb/config"
	"github.com/danish45007/go-dynamodb/internal/repository/adapter"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

func (r *Router) SetRouter(repository adapter.Interface) *chi.Mux {
	r.setConfigRouters()
	r.RouterHealth(repository)
	r.ProductRouter(repository)

	return r.router
}
func (r *Router) setConfigRouters() {
	r.EnableCORS()
	r.EnableRecoverer()
	r.EnableRequestID()
	r.EnableRealIP()
	r.EnableLogger()
	r.EnableTimeout()
}

func (r *Router) RouterHealth(repository adapter.Interface) {

}

func (r *Router) ProductRouter(repository adapter.Interface) {

}

func (r *Router) EnableCORS() *Router {
	r.router.Use(r.config.Cors)
	return r
}

func (r *Router) EnableTimeout() *Router {
	r.router.Use(middleware.Timeout(r.config.GetTimeOut()))
	return r
}
func (r *Router) EnableLogger() *Router {
	r.router.Use(middleware.Logger)
	return r
}

func (r *Router) EnableRecoverer() *Router {
	r.router.Use(middleware.Recoverer)
	return r
}

func (r *Router) EnableRequestID() *Router {
	r.router.Use(middleware.RequestID)
	return r
}

func (r *Router) EnableRealIP() *Router {
	r.router.Use(middleware.RealIP)
	return r
}
