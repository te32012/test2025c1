package application

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"test2025c1/internal/network"
)



func Run() {
	r, err := network.NewServer(fmt.Sprintf("%s:%s", os.Getenv("HOST_APPLICATION"), os.Getenv("PORT_APPLICATION")), 
	fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), 
	os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_BASE")))
	if err != nil {
		panic(err)
	}
	go func() {
		r.Run()
	}()

	slog.Info("success starting application")

	exit := make(chan os.Signal, 2)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	slog.Info("application has been shut down")

}