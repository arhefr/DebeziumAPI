package main

import (
	"context"
	"debez/internal/config"
	repository "debez/internal/repository/postgres"
	"debez/internal/service"
	grpcV1 "debez/internal/transport/grpc/v1"
	httpV1 "debez/internal/transport/http/v1"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"debez/pkg/logger"
	"debez/pkg/postgrespool"

	"go.uber.org/zap/zapcore"
)

func main() {
	// Flags
	envPath := flag.String("env", "config/.env", "path to the environment file")
	flag.Parse()

	// Config
	cfg, err := config.ParseConfig(*envPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Logger
	log, err := logger.NewLogger(zapcore.DebugLevel)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Conn to postgres
	db, err := postgrespool.New(context.Background(), cfg.PostgresConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Repository
	up := repository.NewUserRepository(db.Pool)

	// Service
	us := service.NewUserService(up)

	// HTTP server
	httpHandler := httpV1.NewHandler(us)
	httpServer := httpV1.NewServer(cfg.HTTPServerConfig.Port)
	if err := httpServer.RegisterHandlers(log, httpHandler); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go func() {
		if err := httpServer.Start(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	// GRPC handler
	grpcHandler := grpcV1.NewHandler(us)
	grpcServer := grpcV1.NewServer(log)
	grpcServer.RegisterServices(grpcHandler)

	go func() {
		if err := grpcServer.Run(cfg.GRPCServerConfig.Port); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	graceSh := make(chan os.Signal, 1)
	signal.Notify(graceSh, os.Interrupt, syscall.SIGTERM)
	<-graceSh

	shCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	httpServer.Stop(shCtx)
	grpcServer.Stop()
	db.Stop()
}

/*
+ транспортная часть
+ слой бизнес логики
+ слой работы с данными
+ модели !!!!!!

время разработку
*не переусложни*
*/

/*
+ 1) дописать grpc api
+ 2) подключить базу данных
3) написать запросы !!!!!!
4) заняться кэшированием !!!!!!
*/
