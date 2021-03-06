package main

import (
	"fmt"
	"time"

	"github.com/hahawangxv/kevin"
	"github.com/hahawangxv/kevin/pkg/server"
	"github.com/hahawangxv/kevin/pkg/server/xecho"
	"github.com/hahawangxv/kevin/pkg/server/xgrpc"
	"github.com/hahawangxv/kevin/pkg/worker"
	"github.com/hahawangxv/kevin/pkg/worker/xcron"
	"github.com/labstack/echo/v4"
)

func main() {
	var app jupiter.Application
	app.Startup()
	app.Serve(startHTTPServer())
	app.Serve(startGRPCServer())
	app.Schedule(startWorker())
	app.Run()
}

func startHTTPServer() server.Server {
	server := xecho.DefaultConfig().Build()
	server.GET("/hello", func(ctx echo.Context) error {
		return ctx.JSON(200, "Gopher Wuhan")
	})
	return server
}

func startGRPCServer() server.Server {
	server := xgrpc.DefaultConfig().Build()
	// helloworld.RegisterGreeterServer(server.Server, new(greeter.Greeter))
	return server
}

func startWorker() worker.Worker {
	cron := xcron.DefaultConfig().Build()
	cron.Schedule(xcron.Every(time.Second*10), xcron.FuncJob(func() error {
		fmt.Println("now: ", time.Now().Local().String())
		return nil
	}))
	return cron
}
