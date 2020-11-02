package main

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/selesy/hnjobs/pkg/server"
	"github.com/selesy/hnjobs/pkg/store"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Started HN Jobs Server")

	svrCfg := server.Config{}

	err := envconfig.Process("hnjobs.server", &svrCfg)
	if err != nil {
		// TODO - Log
		os.Exit(1)
	}

	svrCfg.Hostname = ""
	svrCfg.Port = 12345
	svrCfg.Reflection = true

	storeCfg := store.Config{}
	err = envconfig.Process("hnjobs.store", &storeCfg)
	if err != nil {
		// TODO - Log
		os.Exit(2)
	}

	server := server.Server{
		Cfg: svrCfg,
	}

	server.Serve()
}
