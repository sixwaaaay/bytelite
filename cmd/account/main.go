/*
 * Copyright (c) 2023 sixwaaaay.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sixwaaaay/sharing/pkg/configs"
	"github.com/sixwaaaay/sharing/pkg/rpc"
	"github.com/sixwaaaay/sharing/pkg/sign"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Config struct {
	ListenOn    string
	UserService rpc.GrpcConfig
	Jwt         sign.JWT
}

var configFile = flag.String("f", "config.yaml", "the config file")

func main() {
	config, err := configs.NewConfig[Config](*configFile)
	if err != nil {
		panic(err)
	}
	e := newServer()
	client, err := rpc.NewUserClient(config.UserService)
	if err != nil {
		panic(err)
	}
	handler := NewAccountHandler(client, config.Jwt)

	handler.Update(e)

	// Start server
	go func() {
		if err := e.Start(config.ListenOn); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func newServer() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return e
}
