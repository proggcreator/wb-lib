package main

import (
	configreader "git.wildberries.ru/finance/go-infrastructure/config-reader/v2"
	logs "git.wildberries.ru/finance/go-infrastructure/elasticlog"
	_ "github.com/lib/pq"
	restful "github.com/proggcreator/wb-lib"
	"github.com/proggcreator/wb-lib/handler"
	"github.com/proggcreator/wb-lib/repository"
	"github.com/proggcreator/wb-lib/service"
	"github.com/sirupsen/logrus"
)

func main() {

	//get config param
	cfg := repository.Config{}
	err := configreader.Read(&cfg, "configs/config.toml")
	if err != nil {
		logrus.Fatalf("error get configs: %s", err.Error())

		return
	}

	wblogger := logs.NewLogger(
		logs.Settings{
			Host:       cfg.ElasticHost,       //string хост (если необходимо указать несколько хостов, то их необходимо указывать через разделитель ";")
			AppName:    cfg.ElasticAppName,    //string *имя приложения в **kebab-case***
			AppVersion: cfg.ElasticAppVersion, //string *версия приложения*
		})
	//create db connection
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		wblogger.WriteFatal("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db, wblogger)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restful.Server)

	if err := srv.Run(cfg.Port, handlers.InitRoutes()); err != nil {
		wblogger.WriteFatal("error occured while running http server: %s", err.Error())

	}

}
