/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package main is the entrypoint for cutoverlinter.
package main

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"

	"audrahub.com/cutover/cmd/config"
)

func main() {
	var cfg config.Config

	if err := cleanenv.ReadConfig("/home/scarolan/cutover/cmd/config/config.yaml", &cfg); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(cfg.Cutover.AppToken)
	fmt.Println(cfg.Cutover.BaseURL)
	fmt.Println(cfg.Cutover.WorkSpace)
	fmt.Println(cfg.Cutover.InstanceName)

	getApplications()
}
