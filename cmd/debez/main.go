package main

import (
	"context"
	"debez/internal/config"
	repository "debez/internal/repository/postgres"
	"debez/internal/service"
	grpcV1 "debez/internal/transport/grpc/v1"
	httpV1 "debez/internal/transport/http/v1"
	"flag"
	"sync"

	"debez/pkg/logger"
	"debez/pkg/postgrespool"

	"go.uber.org/zap/zapcore"
)

func main() {
	envPath := flag.String("env", "config/.env", "path to the environment file")
	flag.Parse()

	cfg, err := config.ParseConfig(*envPath)
	if err != nil {
		panic(err)
	}

	pp, err := postgrespool.NewPool(context.TODO(), cfg.PostgresConfig)
	if err != nil {
		// panic(err)
	}

	log, err := logger.NewLogger(zapcore.DebugLevel)
	if err != nil {
		panic(err)
	}

	up := repository.NewUserRepository(pp)
	us := service.NewUserService(up)

	var wg sync.WaitGroup

	wg.Go(func() {

		httpHandler := httpV1.NewHandler(us)

		httpServer := httpV1.NewServer(cfg.HTTPServerConfig.Port)

		if err := httpServer.RegisterHandlers(log, httpHandler); err != nil {
			panic(err)
		}

		if err := httpServer.Start(); err != nil {
			panic(err)
		}
	})

	wg.Go(func() {

		grpcHandler := grpcV1.NewHandler(us)

		grpcServer := grpcV1.NewServer(log)

		grpcServer.RegisterServices(grpcHandler)

		if err := grpcServer.Run(cfg.GRPCServerConfig.Port); err != nil {
			panic(err)
		}
	})

	wg.Wait()
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
