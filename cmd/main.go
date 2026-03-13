package main

import (
	// "fmt"
	"log"
	"time"
	"github.com/fc637/go-task-api/internal/handlers"
	// "github.com/fc637/go-task-api/internal/taskstore"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/valyala/fasthttp"
)

var (
	SBIServer *fasthttp.Server
	Concurrency                int    = 1000000
	MAX_TRY                    int    = 5
	Port                       int    =8000
	IPv4Addr                   string    ="localhost"
)
// var 
func main() {
	log.Print("Entering Main \n ")
	server := CreateTaskProducer()
	ServerRun(server, 0)
}


func CreateTaskProducer() *fiber.App {
	log.Print("Entering Task Producer \n ")
	defer log.Print("Exiting Task Producer \n")
	var fiberConf fiber.Config
	fiberConf.Concurrency = 10000
	fiberConf.EnablePrintRoutes = true
	router := fiber.New(fiberConf)
	router.Use(logger.New())
	handlers.AddServices(router)
	return router
}

func ServerRun(arg interface{}, wid int) {
	log.Print("Entering ServerRun \n")
	defer log.Print("Exiting ServerRun\n ")

	router := arg.(*fiber.App)

	SBIServer = &fasthttp.Server{
		Handler:      router.Handler(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	var exitError = true
	var addr string

	for maxTry := MAX_TRY; maxTry > 0; maxTry-- {
		if Port > 0 {
			Port := strconv.Itoa(Port)
			addr = IPv4Addr + ":" + Port

			err := SBIServer.ListenAndServe(addr)
			if err != nil {
				log.Fatalf("Server Run Error EM: %v", err)
				exitError = true
				continue
			}
			exitError = false
			break
		} else {
			log.Fatalf("Server listen Error EM: HTTP port not set \n")
			exitError = true
			break
		}
	}

	if exitError {
		log.Fatalf("Error listening \n")
	}
}