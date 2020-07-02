package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Delisa-sama/logger"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"grpc-boilerplate/api/v1"
	"grpc-boilerplate/config"
	"grpc-boilerplate/controllers"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic("failed to get application configuration: " + err.Error())
	}

	loggerOut := os.Stdout

	if len(cfg.LogPath) <= 0 {
		loggerOut, err = os.OpenFile(cfg.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic("Failed to open log file to write")
		}
	}
	defer loggerOut.Close()

	logger.Init(
		logger.Output(loggerOut),
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
	if err := db.Ping(); err != nil {
		logger.Fatal(err)
	}

	logger.Info("db connected")

	s := grpc.NewServer()
	v1.RegisterEchoServer(s, &controllers.EchoController{})
	v1.RegisterMemoServer(s, &controllers.MemoController{DB: db})

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		if err := s.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			logger.Fatal(err)
		}
	}()

	<-shutdown
	logger.Info("Service stopping")

	s.GracefulStop()

	logger.Info("Service stopped")
}
