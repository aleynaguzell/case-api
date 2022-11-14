package api

import (
	"case-api/api/handler"
	"case-api/pkg/config"
	"case-api/pkg/logger"
	"case-api/pkg/mongo"
	"case-api/services"
	"case-api/storage/cache"
	"case-api/storage/repository"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	defaultEnv = "dev"
)

type IApp interface {
	Start()
}

type app struct {
}

var App *app

func Init() IApp {
	App = &app{}
	return App
}

func (a *app) Start() {
	env := getEnvironment("APP_ENVIRONMENT", defaultEnv)
	config.Setup(".", env)

	mongoClient, err := mongo.Init()

	if err != nil {
		fmt.Printf("could not configure mongoClient: %v", err)
		os.Exit(1)
	}

	logger.Init()
	log.Printf("logger is active now")

	memory := cache.New()
	recordRepository := repository.NewRecordsRepository(mongoClient)
	recordService := services.NewRecordService(recordRepository)
	recordHandler := handler.NewRecordHandler(*recordService)

	memoryService := services.NewMemoryService(memory)
	memoryHandler := handler.NewMemoryHandler(*memoryService)

	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.HandleFunc("/", HealthCheck)
	http.HandleFunc("/in-memory/", memoryHandler.Set)
	http.HandleFunc("/in-memory", memoryHandler.Get)
	http.HandleFunc("/records", recordHandler.Get)

	setHttpClient()
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}

func getEnvironment(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func setHttpClient() {

	port := os.Getenv("PORT")
	if port == "" {
		port = config.GetConfig().Server.Port
	}
	httpServer := &http.Server{
		Addr: ":" + port,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("HTTP server ListenAndServe: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	<-signalChan
	log.Printf("os.Interrupt - shutting down...\n")

	go func() {
		<-signalChan
		log.Fatal("os.Kill - terminating...\n")
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("error handled: %v\n", err)
		defer os.Exit(1)
		return
	} else {
		log.Printf("stopped\n")
	}

	defer os.Exit(0)
}
