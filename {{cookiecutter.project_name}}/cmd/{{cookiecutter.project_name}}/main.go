package main

import (
	"log"

	"{{cookiecutter.module_path}}/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
