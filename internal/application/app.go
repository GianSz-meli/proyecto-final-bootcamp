package application

import (
	"ProyectoFinal/internal/application/di"
	"ProyectoFinal/internal/application/loader"
	"ProyectoFinal/internal/application/router"
	"ProyectoFinal/internal/handler"
	"ProyectoFinal/internal/repository"
	"ProyectoFinal/internal/service"
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
	sellerDB, err := factory.NewSellerLoader().Load()
	if err != nil {
		panic(err)
	}
	warehouseDB, err := factory.NewWarehouseLoader().Load()
	if err != nil {
		panic(err)
	}
	buyerDB, err := factory.NewBuyerLoader().Load()
	if err != nil {
		panic(err)
	}

	warehouseHandler := di.GetWarehouseHandler(warehouseDB)
	buyerHandler := di.GetBuyerHandler(buyerDB)

	// Load sections
	sections, err := factory.NewSectionLoader().Load()
	if err != nil {
		panic(err)
	}

	sectionHandler := di.GetSectionHandler(sections)

	repoSeller := repository.NewSellerRepository(sellerDB)
	srvSeller := service.NewSellerService(repoSeller)
	ctrSeller := handler.NewSellerHandler(srvSeller)

	rt.Route("/api/v1", func(r chi.Router) {

		r.Mount("/sections", router.SectionRoutes(sectionHandler))

		r.Mount("/seller", router.SellerRoutes(ctrSeller))
		r.Mount("/warehouses", router.GetWarehouseRouter(warehouseHandler))
		r.Mount("/buyers", router.GetBuyerRouter(buyerHandler))
	})

	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
