package exporter

import (
	"sync"

	"github.com/OpenNebula/one/src/oca/go/src/goca/schemas/datastore"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	dsStatusOnce sync.Once
	dsTotalOnce  sync.Once
	dsUsedOnce   sync.Once
	dsFreeOnce   sync.Once
)
var (
	dsStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "opennebula_monitoring_datastore_status",
			Help: "OpenNebula Host Status (0 = ON, 1 = OFF)",
		},
		[]string{"hostname"},
	)

	dsTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "opennebula_monitoring_datastore_total",
			Help: "Total datastore capacity",
		},
		[]string{"hostname"},
	)
	dsFree = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "opennebula_monitoring_datastore_Free",
			Help: "Datastore Free space MB",
		},
		[]string{"hostname"},
	)
	dsUsed = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "opennebula_monitoring_datastore_used",
			Help: "Datastroe Used space MB",
		},
		[]string{"hostname"},
	)
)

func RegisterDataStoreMetrics(datastores *datastore.Pool) {

	for _, ds := range datastores.Datastores {
		dsStatus.WithLabelValues(ds.Name).Set(float64(ds.StateRaw))
		dsTotal.WithLabelValues(ds.Name).Set(float64(ds.TotalMB))
		dsFree.WithLabelValues(ds.Name).Set(float64(ds.FreeMB))
		dsUsed.WithLabelValues(ds.Name).Set(float64(ds.UsedMB))

	}
	dsStatusOnce.Do(func() {
		prometheus.MustRegister(dsStatus)

	})
	dsTotalOnce.Do(func() {
		prometheus.MustRegister(dsTotal)
	})
	dsFreeOnce.Do(func() {
		prometheus.MustRegister(dsFree)

	})
	dsUsedOnce.Do(func() {
		prometheus.MustRegister(dsUsed)

	})

}
