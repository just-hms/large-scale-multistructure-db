package main

import (
	"fmt"

	"github.com/just-hms/large-scale-multistructure-db/be/config"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/app"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println("error retriving config file: ", err)
		return
	}
	app.Run(cfg)
}
