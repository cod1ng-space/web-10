package main

import (
	"flag"
	"log"

	"github.com/cod1ng-space/web-10/internal/query/api"
	"github.com/cod1ng-space/web-10/internal/query/config"
	"github.com/cod1ng-space/web-10/internal/query/provider"
	"github.com/cod1ng-space/web-10/internal/query/usecase"
	_ "github.com/lib/pq"
)

func main() {
	// Считываем аргументы командной строки
	configPath := flag.String("config-path", "./configs/query.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase(prv, cfg.Usecase.StartMessage, cfg.Usecase.EndMessage)
	srv := api.NewServer(cfg.IP, cfg.Port, cfg.API.MaxMessageSize, use)

	srv.Run()
}
