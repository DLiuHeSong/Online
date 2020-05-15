package api

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.elastic.co/apm/module/apmechov4"

	"Online/handler/middle"
)

func Start() {
	flag.Parse()
	// New a echo example
	e := echo.New()
	// open a log file
	fileName := path.Join("./Online.log")
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Errorf("failed create log file, err is %s", err))
	}
	// user a echo log format
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logFile,
		Format: "${time_rfc3339} ${status} ${method} ${uri} ${latency_human} ${remote_ip} ${bytes_in} ${bytes_out} ${user_agent} ${referer}\n",
	}))
	// validate request_id middleware ...
	e.Use(apmechov4.Middleware())
	e.Use(middleware.RequestID())
	e.Use(middle.Validate)
	// 注册router
	//init the routers
	InitRouter(e)
	// start the project
	go func() {
		e.Logger.Fatal(e.Start("127.0.0.1:8080"))
	}()

	// 信号
	// select the signal
	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	s := <-signalCh
	switch s {
	case syscall.SIGTERM, syscall.SIGINT:
		c, _ := context.WithTimeout(context.Background(), time.Second*5)
		_ = e.Shutdown(c)
		log.Println("receive signal is ", s)
		os.Exit(0)
	}
}
