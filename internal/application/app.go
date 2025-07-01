package application

import (
	"ProyectoFinal/internal/application/loader"
	handler "ProyectoFinal/internal/handler"
	repository "ProyectoFinal/internal/repository/section"
	service "ProyectoFinal/internal/service/section"
	"net/http"

	"github.com/go-chi/chi/v5"
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

	// Load sellers
	_, err = factory.NewSellerLoader().Load()
	if err != nil {
		panic(err)
	}

	sections, err := factory.NewSectionLoader().Load()
	if err != nil {
		panic(err)
	}

	sectionRepo := repository.NewSectionMap(sections)
	sectionService := service.NewSectionDefault(sectionRepo)
	sectionHandler := handler.NewSectionDefault(sectionService)

	rt.Route("/api/v1", func(r chi.Router) {
		r.Mount("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong"))
		})) // TODO: remove this

		// Section routes
		r.Route("/sections", func(r chi.Router) {
			r.Get("/", sectionHandler.GetAll())
			r.Get("/{id}", sectionHandler.GetByID())
			r.Post("/", sectionHandler.Create())
			r.Put("/{id}", sectionHandler.Update())
			r.Delete("/{id}", sectionHandler.Delete())
		})
	})

	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
