package application

import (
	"ProyectoFinal/internal/application/config"
	"ProyectoFinal/internal/application/di"
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

	dbConn := config.InitDB()
	// Dependency injection
	sectionHandler := di.GetSectionHandler(dbConn)
	// TODO: Actualizar los siguientes handlers para usar MySQL si es necesario
	// sellerHandler := di.GetSellerHandler(...)
	// warehouseHandler := di.GetWarehouseHandler(...)
	// buyerHandler := di.GetBuyerHandler(...)
	// employeeHandler := di.GetEmployeeHandler(...)
	// productHandler := di.GetProductsHandler(...)

	//Middlewares
	rt.Use(middleware.Logger)

	rt.Route("/api/v1", func(r chi.Router) {
		r.Mount("/sections", router.GetSectionRouter(sectionHandler))
		// r.Mount("/sellers", router.GetSellerRouter(sellerHandler))
		// r.Mount("/employees", router.EmployeeRoutes(employeeHandler))
		// r.Mount("/warehouses", router.GetWarehouseRouter(warehouseHandler))
		// r.Mount("/products", router.ProductRoutes(productHandler))
		// r.Mount("/buyers", router.GetBuyerRouter(buyerHandler))
	})

	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
