package main

import (
	"fmt"

	"github.com/erjvarela/trollgo/internal/config"
)

func main() {
	Cfg, err := config.InitEnviromentVariables()
	if err != nil {
		fmt.Println("Error loading environment variables:", err)
		panic(err)
	}
	fmt.Printf("%v", Cfg.SshDefaultUser)
}
