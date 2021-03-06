package main

import (
	"github.com/DiscoreMe/SecureCloud/config"
	"github.com/DiscoreMe/SecureCloud/internal"
)

func main() {
	cfg := config.NewConfig()

	serv := internal.NewServer(internal.ServerConfig{
		ValidToken: cfg.Token,
	})

	serv.SetupAPI()

	if err := serv.EnableLocalStorage(); err != nil {
		panic(err)
	}

	if cfg.IsSupportDrive {
		if err := serv.EnableDriveStorage(); err != nil {
			if err := serv.UpdateDriveToken(); err != nil {
				panic(err)
			}
		}
	}

	if err := serv.Listen("127.0.0.1:7777"); err != nil {
		panic(err)
	}
}
