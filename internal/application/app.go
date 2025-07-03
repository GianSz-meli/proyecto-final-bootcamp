package application

import (
	"ProyectoFinal/internal/application/di"
	"ProyectoFinal/internal/application/loader"
	"ProyectoFinal/internal/application/router"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

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

	//Load warehouse
	warehouseDB, err := factory.NewWarehouseLoader().Load()
	if err != nil {
		panic(err)
	}

	//Load buyer
	buyerDB, err := factory.NewBuyerLoader().Load()
	if err != nil {
		panic(err)
	}

	// Load sections
	sections, err := factory.NewSectionLoader().Load()
	if err != nil {
		panic(err)
	}

	// Load employee data
	employeeDB, err := factory.NewEmployeeLoader().Load()
	if err != nil {
		panic(err)
	}

	// Dependency injection
	sellerHandler := di.GetSellerHandler(sellerDB)
	warehouseHandler := di.GetWarehouseHandler(warehouseDB)
	sectionHandler := di.GetSectionHandler(sections)
	buyerHandler := di.GetBuyerHandler(buyerDB)
	employeeHandler := di.GetEmployeeHandler(employeeDB)

	//Middlewares
	rt.Use(middleware.Logger)

	rt.Route("/api/v1", func(r chi.Router) {
		r.Mount("/sections", router.GetSectionRouter(sectionHandler))
		r.Mount("/sellers", router.GetSellerRouter(sellerHandler))
		r.Mount("/employees", router.EmployeeRoutes(employeeHandler))
		r.Mount("/warehouses", router.GetWarehouseRouter(warehouseHandler))
		r.Mount("/buyers", router.GetBuyerRouter(buyerHandler))
	})

	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
