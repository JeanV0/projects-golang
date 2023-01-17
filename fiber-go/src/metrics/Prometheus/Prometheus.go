package prometheus

import (
	dotenv "marketplace/config/services/Dotenv"
	"marketplace/config/services/db"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/dlmiddlecote/sqlstats"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

func configPrometheus() {
	collector := sqlstats.NewStatsCollector(dotenv.MyEnvironmentApp.Database_DbName, db.DbInstance.DB())
	prometheus.MustRegister(collector)
}

func GetPrometheusFiber(FiberApp *fiber.App) {
	configPrometheus()
	prometheus := fiberprometheus.New("pigz_metrics")
	prometheus.RegisterAt(FiberApp, "/metrics/fiber")
	FiberApp.Use(prometheus.Middleware)

}
