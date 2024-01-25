package infra

import (
	"context"
	infra_http "cuddly-octo-palm-tree/infra/http"

	// infra_kafka "go_microservice/infra/kafka"

	log "github.com/sirupsen/logrus"
)

type RuServer struct {
	HTTPServer infra_http.Server
	// Consumer   infra_kafka.Consumer
}

type BiServer struct {
	HTTPServer infra_http.Server
	// Consumer   infra_kafka.Consumer
}

type ProjectStatusServer struct {
	HTTPServer infra_http.Server
	// Consumer   infra_kafka.Consumer
}

type TrServer struct {
	HTTPServer infra_http.Server
	// Consumer   infra_kafka.Consumer
}

// NewProductServer factory
func NewRuServer(httpServer infra_http.Server) *RuServer {
	return &RuServer{
		HTTPServer: httpServer,
	}
}

// Run server
func (s *RuServer) Run() error {

	go func() {
		err := s.HTTPServer.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// go func() {
	// 	err := s.Consumer.Run()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	return nil
}

// GracefulStop server
func (s *RuServer) GracefulStop(ctx context.Context, done chan bool) {
	err := s.HTTPServer.GracefulStop(ctx)
	if err != nil {
		log.Error(err)
	}

	// err = s.Consumer.GracefulStop()
	// if err != nil {
	// 	log.Error(err)
	// }

	log.Info("gracefully shutdowned")
	done <- true
}

// NewProductServer factory
func NewBiServer(httpServer infra_http.Server) *BiServer {
	return &BiServer{
		HTTPServer: httpServer,
	}
}

// Run server
func (s *BiServer) Run() error {

	go func() {
		err := s.HTTPServer.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// go func() {
	// 	err := s.Consumer.Run()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	return nil
}

// GracefulStop server
func (s *BiServer) GracefulStop(ctx context.Context, done chan bool) {
	err := s.HTTPServer.GracefulStop(ctx)
	if err != nil {
		log.Error(err)
	}

	// err = s.Consumer.GracefulStop()
	// if err != nil {
	// 	log.Error(err)
	// }

	log.Info("gracefully shutdowned")
	done <- true
}

// NewProductServer factory
func NewProjectStatusServer(httpServer infra_http.Server) *ProjectStatusServer {
	return &ProjectStatusServer{
		HTTPServer: httpServer,
	}
}

// Run server
func (s *ProjectStatusServer) Run() error {

	go func() {
		err := s.HTTPServer.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// go func() {
	// 	err := s.Consumer.Run()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	return nil
}

// GracefulStop server
func (s *ProjectStatusServer) GracefulStop(ctx context.Context, done chan bool) {
	err := s.HTTPServer.GracefulStop(ctx)
	if err != nil {
		log.Error(err)
	}

	// err = s.Consumer.GracefulStop()
	// if err != nil {
	// 	log.Error(err)
	// }

	log.Info("gracefully shutdowned")
	done <- true
}

// NewProductServer factory
func NewTrServer(httpServer infra_http.Server) *TrServer {
	return &TrServer{
		HTTPServer: httpServer,
	}
}

// Run server
func (s *TrServer) Run() error {

	go func() {
		err := s.HTTPServer.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// go func() {
	// 	err := s.Consumer.Run()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	return nil
}

// GracefulStop server
func (s *TrServer) GracefulStop(ctx context.Context, done chan bool) {
	err := s.HTTPServer.GracefulStop(ctx)
	if err != nil {
		log.Error(err)
	}

	// err = s.Consumer.GracefulStop()
	// if err != nil {
	// 	log.Error(err)
	// }

	log.Info("gracefully shutdowned")
	done <- true
}
