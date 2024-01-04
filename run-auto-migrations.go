// Copyright 2023 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gocondor/core"
	"github.com/harranali/gocondor-todo-app/models"
)

func RunAutoMigrations() {
	db := core.ResolveGorm()
	//##############################
	//# Models auto migration  #####
	//##############################

	// Add auto migrations for your models here...
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Todo{})
}
