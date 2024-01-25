package backlogimpl

import (
	"context"
	"cuddly-octo-palm-tree/config"
	"cuddly-octo-palm-tree/infra/http/middleware"
	"io"
	"net/http"

	infra_http "cuddly-octo-palm-tree/infra/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type BiServer struct {
	App    string
	Port   string
	Engine *gin.Engine
	Router *Router
	svr    *http.Server
}

// GracefulStop implements http.Server.
func (*BiServer) GracefulStop(ctx context.Context) error {
	panic("unimplemented")
}

// RegisterRoutes implements http.Server.
func (s *BiServer) RegisterRoutes() {
	s.Engine.Use(middleware.CorsMiddleware())
	apiGroup := s.Engine.Group("/api/backlogimpl")
	{
		apiGroup.GET("/users", s.Router.GetUsers)
		apiGroup.GET("/orders", s.Router.GetOrder)
		apiGroup.GET("/items", s.Router.GetOrderItems)
		apiGroup.GET("/products", s.Router.GetProducts)
		apiGroup.GET("/users/:id", s.Router.GetUserById)
		apiGroup.DELETE("/users/:id", s.Router.DeleteUser)
	}
}

// Start server
func (s *BiServer) Run() error {
	s.RegisterRoutes()
	addr := ":" + s.Port
	s.svr = &http.Server{
		Addr:    addr,
		Handler: infra_http.NewOtelHandler(s.Engine, s.App+"_http"),
	}
	log.Infoln("http server listening on ", addr)
	err := s.svr.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func NewEngine(conf *config.Config) *gin.Engine {
	gin.SetMode(conf.GinMode)
	if conf.GinMode == "release" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}
	gin.DefaultWriter = io.Writer(conf.Logger.Writer)

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.LogMiddleware(conf.Logger.ContextLogger))

	return engine
}

// Factory of product server
func NewBiServer(conf *config.Config, engine *gin.Engine, router *Router) *BiServer {
	return &BiServer{
		App:    conf.App,
		Port:   conf.HTTPPort,
		Engine: engine,
		Router: router,
	}
}
