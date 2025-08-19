package application

import (
	"ProyectoFinal/internal/application/config"
	"ProyectoFinal/internal/application/di"
	"ProyectoFinal/internal/application/router"
	"ProyectoFinal/pkg/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ConfigServerChi struct {
	ServerAddress string
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
	}

	return &ServerChi{
		serverAddress: defaultConfig.ServerAddress,
	}
}

func (a *ServerChi) Run() (err error) {
	rt := chi.NewRouter()
	sqlDB := config.InitDB()

	// Dependency injection
	sellerHandler := di.GetSellerHandler(sqlDB)
	warehouseHandler := di.GetWarehouseHandler(sqlDB)
	sectionHandler := di.GetSectionHandler(sqlDB)
	buyerHandler := di.GetBuyerHandler(sqlDB)
	employeeHandler := di.GetEmployeeHandler(sqlDB)
	productHandler := di.GetProductsHandler(sqlDB)
	productBatchHandler := di.GetProductBatchHandler(sqlDB)
	carrierHandler := di.GetCarrierHandler(sqlDB)
	localityHandler := di.GetLocalityHandler(sqlDB)
	inboundOrderHandler := di.GetInboundOrderHandler(sqlDB)
	purchaseOrderHandler := di.GetPurchaseOrderHandler(sqlDB)
	productRecordHandler := di.GetProductRecordHandler(sqlDB)

	//Middlewares
	rt.Use(middlewares.Logger(sqlDB))

	rt.Route("/api/v1", func(r chi.Router) {
		r.Mount("/sections", router.GetSectionRouter(sectionHandler, productBatchHandler))
		r.Mount("/productBatches", router.GetProductBatchRouter(productBatchHandler))
		r.Mount("/sellers", router.GetSellerRouter(sellerHandler))
		r.Mount("/employees", router.EmployeeRoutes(employeeHandler, inboundOrderHandler))
		r.Mount("/warehouses", router.GetWarehouseRouter(warehouseHandler))
		r.Mount("/products", router.ProductRoutes(productHandler, productRecordHandler))
		r.Mount("/buyers", router.GetBuyerRouter(buyerHandler))
		r.Mount("/productRecords", router.GetProductRecordRouter(productRecordHandler))
		r.Mount("/localities", router.GetLocalityRouter(localityHandler))
		r.Mount("/carriers", router.GetCarrierRouter(carrierHandler))
		r.Mount("/inboundOrders", router.InboundOrderRoutes(inboundOrderHandler))
		r.Mount("/purchaseOrders", router.GetPurchaseOrderRouter(purchaseOrderHandler))
	})

	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
