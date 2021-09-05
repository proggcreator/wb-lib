package main

import (
	configreader "git.wildberries.ru/finance/go-infrastructure/config-reader/v2"
	_ "github.com/lib/pq"
	restful "github.com/proggcreator/wb-lib"
	"github.com/proggcreator/wb-lib/handler"
	"github.com/proggcreator/wb-lib/repository"
	"github.com/proggcreator/wb-lib/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	//set loger format

	//get config param
	config := repository.Config{}
	err := configreader.Read(&config, "../configs/config.toml")
	if err != nil {
		logrus.Fatalf("error get configs: %s", err.Error())
		return
	}
	//create db connection
	db, err := repository.NewPostgresDB(config)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restful.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}
