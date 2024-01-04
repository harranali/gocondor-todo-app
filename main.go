// Copyright 2023 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"path"

	"github.com/gocondor/core"
	"github.com/gocondor/core/env"
	"github.com/gocondor/core/logger"
	"github.com/harranali/gocondor-todo-app/config"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

// The main function
func main() {
	app := core.New()
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatal("error getting current working dir")
	}
	app.SetBasePath(basePath)
	app.MakeDirs("logs", "storage", "storage/sqlite", "tls")
	// Handle the reading of the .env file
	if config.GetEnvFileConfig().UseDotEnvFile {
		envVars, err := godotenv.Read(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		env.SetEnvVars(envVars)
	}
	// Handle the logs
	app.SetLogsDriver(&logger.LogFileDriver{
		FilePath: path.Join(basePath, "logs/app.log"),
	})
	app.SetRequestConfig(config.GetRequestConfig())
	app.SetGormConfig(config.GetGormConfig())
	app.SetCacheConfig(config.GetCacheConfig())
	app.Bootstrap()
	registerGlobalMiddlewares()
	registerRoutes()
	registerEvents()
	if config.GetGormConfig().EnableGorm == true {
		RunAutoMigrations()
	}
	app.Run(httprouter.New())
}
