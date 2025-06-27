package application

import (
	"ProyectoFinal/internal/application/loader"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ConfigServerChi struct {
	ServerAddress  string
	LoaderFilePath map[string]string
}

type ServerChi struct {
	serverAddress  string
	loaderFilePath map[string]string
}

func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if len(cfg.LoaderFilePath) != 0 {
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

	factory := loader.NewLoaderFactory(a.loaderFilePath)

	sellersDB, err := factory.NewSellerLoader().Load()

	if err != nil {
		panic(err)
	}

	// fmt.Println(a.loaderFilePath)

	fmt.Println(sellersDB)
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
