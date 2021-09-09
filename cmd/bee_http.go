package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/nexclipper/curlbee/pkg/policy"
)

type HttpBee struct {
	handler *BeeHandler
	params  string
}

func (c *HttpBee) Run(cfg []policy.BeePolicy) error {
	route := echo.New()
	route.Use(middleware.Logger())
	route.Use(middleware.Recover())

	route.POST("/call", c.handler.Call)

	go func() {
		if err := route.Start(":3001"); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	select {
	case <-quit:
		log.Println("os interrupt")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := route.Shutdown(ctx); err != nil {
			return err
		}
	}

	return nil
}

type BeeHandler struct{}

func (b *BeeHandler) Call(ctx echo.Context) error {
	log.Println("Call...start")
	return nil
}
