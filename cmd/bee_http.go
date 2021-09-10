package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/nexclipper/curlbee/pkg/client"
	"github.com/nexclipper/curlbee/pkg/config"
	"github.com/nexclipper/curlbee/pkg/util"
)

type HttpBee struct {
	handler *BeeHandler
}

func (c *HttpBee) Run(cfg *config.BeeConfig) error {

	if cfg.Title == "" {
		return fmt.Errorf("title is empty")
	}

	if cfg.Port == 0 {
		return fmt.Errorf("port is empty")
	}

	route := echo.New()
	route.Use(middleware.Logger())
	route.Use(middleware.Recover())

	c.handler.RegistPolicy(cfg.Policies)

	pathName := fmt.Sprintf("/%s", strings.ReplaceAll(strings.TrimSpace(strings.ToLower(cfg.Title)), " ", ""))

	// make it POST to receive parameters
	route.POST(pathName, c.handler.Action)

	go func() {
		if err := route.Start(fmt.Sprintf(":%d", cfg.Port)); err != nil {
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

type BeeHandler struct {
	Policies []config.BeePolicy
}

func (b *BeeHandler) RegistPolicy(p []config.BeePolicy) {
	b.Policies = p
}

func (b *BeeHandler) Action(ctx echo.Context) error {
	var result string
	if body, err := ioutil.ReadAll(ctx.Request().Body); err != nil {
		return err
	} else {
		respBuf := make(map[string]string)
		params := string(body)
		parameters := util.SplitParameter(params)
		for _, p := range b.Policies {
			var name, respBody string
			p.VariableMatching(parameters)
			err := client.Request(&p, &name, &respBody)
			if err != nil {
				return err
			} else {
				respBuf[name] = respBody
			}

		}

		for k, v := range respBuf {
			result += fmt.Sprintf("%s\n\n%s\n\n", k, v)
		}
	}

	return ctx.String(http.StatusOK, result)
}
