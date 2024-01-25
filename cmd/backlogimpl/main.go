package main

import (
	"context"
	"cuddly-octo-palm-tree/config"
	"cuddly-octo-palm-tree/infra"
	"cuddly-octo-palm-tree/infra/db/pgserver"
	"os"
	"os/signal"
	"syscall"
	"time"

	bi_http "cuddly-octo-palm-tree/infra/http/backlogimpl"
	bi_repo "cuddly-octo-palm-tree/repo/backlogimpl"
	bi_service "cuddly-octo-palm-tree/service/backlogimpl"

	log "github.com/sirupsen/logrus"
)

func main() {
	migrator, err := InitializeMigrator()
	if err != nil {
		log.Fatal(err)
	}
	if err := migrator.Migrate(); err != nil {
		log.Fatal(err)
	}

	server, err := InitializeServer()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		err := server.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// catch shutdown
	done := make(chan bool, 1)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		// graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.GracefulStop(ctx, done)
	}()

	// wait for graceful shutdown
	<-done
}

func InitializeServer() (*infra.BiServer, error) {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	gormDB, err := pgserver.NewDatabaseConnection(conf)
	if err != nil {
		log.Fatal(err)
	}
	// producer, err := kafka.NewProducer(conf)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	engine := bi_http.NewEngine(conf)

	biRepository := bi_repo.NewBiRepository(gormDB)
	biService := bi_service.NewBiService(conf, biRepository)
	// producer, err := broker.NewProducer(conf)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	biRouter := bi_http.NewRouter(biService)

	server := bi_http.NewBiServer(conf, engine, biRouter)
	// aConsumerRouter := a_broker.NewRouter(aService)
	// consumer := a_broker.NewAConsumer(conf, aConsumerRouter)
	biServer := infra.NewBiServer(server)

	return biServer, nil
}

func InitializeMigrator() (*pgserver.Migrator, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	gormDB, err := pgserver.NewDatabaseConnection(configConfig)
	if err != nil {
		return nil, err
	}
	migrator := pgserver.NewMigrator(gormDB)
	return migrator, nil
}
