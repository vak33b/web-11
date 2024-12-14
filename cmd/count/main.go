package main

import (
	"flag"
	"github.com/ValeryBMSTU/web-11/internal/count/api"
	"github.com/ValeryBMSTU/web-11/internal/count/config"
	"github.com/ValeryBMSTU/web-11/internal/count/provider"
	"github.com/ValeryBMSTU/web-11/internal/count/usecase"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Считываем аргументы командной строки
	configPath := flag.String("config-path", "configs/config.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase(cfg.Usecase.DefaultMessageCount, prv)
	srv := api.NewServer(cfg.IP, cfg.Port, cfg.API.MaxMessageSize, use)

	srv.Run()
}
