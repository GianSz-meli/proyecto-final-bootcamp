package application

import (
	"ProyectoFinal/docs/db"
	"ProyectoFinal/internal/application/config"
	"ProyectoFinal/internal/application/di"
	"ProyectoFinal/internal/application/router"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	database := db.LoadDB(a.loaderFilePath)
	dbConn := config.InitDB()

	// Dependency injection
	sellerHandler := di.GetSellerHandler(database.Seller)
	warehouseHandler := di.GetWarehouseHandler(database.Warehouse)
	sectionHandler := di.GetSectionHandler(dbConn)
	buyerHandler := di.GetBuyerHandler(database.Buyer)
	employeeHandler := di.GetEmployeeHandler(database.Employee)
	productHandler := di.GetProductsHandler(database.Product)

	//Middlewares
	rt.Use(middleware.Logger)

	rt.Route("/api/v1", func(r chi.Router) {
		r.Mount("/sections", router.GetSectionRouter(sectionHandler))
		r.Mount("/sellers", router.GetSellerRouter(sellerHandler))
		r.Mount("/employees", router.EmployeeRoutes(employeeHandler))
		r.Mount("/warehouses", router.GetWarehouseRouter(warehouseHandler))
		r.Mount("/products", router.ProductRoutes(productHandler))
		r.Mount("/buyers", router.GetBuyerRouter(buyerHandler))
	})

	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
