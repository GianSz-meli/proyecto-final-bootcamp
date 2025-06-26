package application

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ConfigServerChi struct {
	ServerAddress  string
	LoaderFilePath string
}

type ServerChi struct {
	serverAddress  string
	loaderFilePath string
}

func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
	}

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

func (a *ServerChi) Run() (err error) {
	rt := chi.NewRouter()
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
