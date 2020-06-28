package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Delisa-sama/logger"
	"google.golang.org/grpc"
	"grpc-boilerplate/api"
	"grpc-boilerplate/config"
	"grpc-boilerplate/controllers"
	"net"
	"os"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic("failed to get application configuration: " + err.Error())
	}

	logFile, err := os.OpenFile(cfg.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Failed to open log file to write")
	}
	defer logFile.Close()

	logger.Init(
		logger.Output(logFile),
		logger.Level(cfg.LogLevel),
	)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	if err != nil {
		logger.Fatal(err)
	}

	cfgString, _ := json.Marshal(cfg)
	logger.Infof("start service with configuration: %s", string(cfgString))

	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Name,
	)

	db, err := sql.Open(cfg.DB.Dialect, dbURI)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("db connected")

	s := grpc.NewServer()
	api.RegisterEchoServer(s, &controllers.EchoController{})
	api.RegisterMemoServer(s, &controllers.MemoController{DB: db})

	if err := s.Serve(lis); err != nil {
		logger.Fatal(err)
	}
}
