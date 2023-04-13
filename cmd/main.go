package main

import (
	"OneLab2/config"
	handler2 "OneLab2/handler"
	service2 "OneLab2/service"
	"OneLab2/storage"
	"OneLab2/storage/postgres"
	"OneLab2/transport/http"
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Fatalln(run())
}

func run() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	db, err := postgres.InitDb(*cfg)
	if err != nil {
		return err
	}
	repo, err := storage.NewStorage(db)
	if err != nil {
		return err
	}
	service, err := service2.NewService(repo)
	if err != nil {
		return err
	}
	handler, err := handler2.NewHandler(service)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	srv, err := http.NewServer(cfg, ctx, handler)
	if err != nil {
		return err
	}
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	go func() {
		log.Println(<-sig)
		cancel()
	}()
	err = srv.Run()
	if err != nil {
		return err
	}
	<-srv.Conn
	return nil
}
