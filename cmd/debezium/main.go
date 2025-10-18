package main

import (
	"debez/internal/config"
	"debez/internal/repository"
	"debez/internal/service"
	v1 "debez/internal/transport/http/v1"
	"flag"
)

func main() {
	envPath := flag.String("env", ".env", "path to the environment file")
	flag.Parse()
	cfg, err := config.ParseConfig(*envPath)
	if err != nil {
		return
	}

	up := repository.NewUserRepository()

	us := service.NewUserService(up)

	server := v1.NewServer(cfg.Port)

	if err := server.Start(); err != nil {
		return
	}
}

/*
транспортная часть
слой бизнес логики
слой работы с данными
модели


время разработку
*не переусложни*
*/

/*
1) дописать grpc api
2) подключить базу данных
3) написать запросы
4) заняться кэшированием
*/
