package routes

import (
	ServerConfig "github.com/Luis-Miguel-BL/go-dynamodb-crud/config"
	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/http/controllers"
	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/repositories"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig().SetTimeout(ServerConfig.GetConfig().Timeout),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouters(emailScoreRepository repositories.EmailScoreRepository, heathRepository repositories.HealthRepository) *chi.Mux {
	r.setConfigsRouters()

	r.RouterEmailScore(emailScoreRepository)
	r.RouterHealth(heathRepository)

	return r.router
}

func (r *Router) setConfigsRouters() {
	r.EnableCORS()
	r.EnableLogger()
	r.EnableTimeout()
	r.EnableRecover()
	r.EnableRequestID()
	r.EnableRealIP()
}

func (r *Router) RouterHealth(repository repositories.HealthRepository) {
	controller := controllers.NewHealthController(repository)

	r.router.Route("/health", func(route chi.Router) {
		route.Get("/", controller.Get)
	})
}

func (r *Router) RouterEmailScore(repository repositories.EmailScoreRepository) {
	controller := controllers.NewEmailScoreController(repository)

	r.router.Route("/email-score", func(route chi.Router) {
		route.Post("/find", controller.FindByEmails)
		route.Get("/by-email/", controller.GetByEmail)
		route.Put("/consolidate-score/", controller.ConsolidateScore)
	})
}

func (r *Router) EnableLogger() *Router {
	r.router.Use(middleware.Logger)
	return r
}

func (r *Router) EnableTimeout() *Router {
	r.router.Use(middleware.Timeout(r.config.GetTimeout()))
	return r
}

func (r *Router) EnableCORS() *Router {
	r.router.Use(r.config.Cors)
	return r
}

func (r *Router) EnableRecover() *Router {
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
