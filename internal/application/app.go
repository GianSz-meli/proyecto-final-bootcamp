package application

import (
	"ProyectoFinal/internal/application/di"
	"ProyectoFinal/internal/application/loader"
	"ProyectoFinal/internal/application/router"
	"ProyectoFinal/internal/handler"
	employeeHandler "ProyectoFinal/internal/handler/employee"
	"ProyectoFinal/internal/repository"
	employeeRepository "ProyectoFinal/internal/repository/employee"
	"ProyectoFinal/internal/service"
	employeeService "ProyectoFinal/internal/service/employee"
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

	sellerDB, err := factory.NewSellerLoader().Load()
	if err != nil {
		panic(err)
	}

	warehouseDB, err := factory.NewWarehouseLoader().Load()
	if err != nil {
		panic(err)
	}

	warehouseHandler := di.GetWarehouseHandler(warehouseDB)

	repoSeller := repository.NewSellerRepository(sellerDB)
	srvSeller := service.NewSellerService(repoSeller)
	ctrSeller := handler.NewSellerHandler(srvSeller)

	employeeDB, err := factory.NewEmployeeLoader().Load()
	if err != nil {
		panic(err)
	}
	repoEmployee := employeeRepository.NewRepository(employeeDB)
	srvEmployee := employeeService.NewService(repoEmployee)
	ctrEmployee := employeeHandler.NewEmployeeHandler(srvEmployee)

	rt.Route("/api/v1", func(r chi.Router) {
		r.Mount("/seller", router.SellerRoutes(ctrSeller))
		r.Mount("/employees", router.EmployeeRoutes(ctrEmployee))
		r.Mount("/warehouses", router.GetWarehouseRouter(warehouseHandler))
	})

	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
