package exporter

import (
	"sync"

	"github.com/OpenNebula/one/src/oca/go/src/goca/schemas/virtualnetwork"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	vNetsTotalOnce  sync.Once
	vNetsLeasesOnce sync.Once
)
var (
	vNetsTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "opennebula_monitoring_vnet_total",
			Help: "Total number of available IPs in vnet pools",
		},
		[]string{"vnet"},
	)

	vNetsLeases = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "opennebula_monitoring_vnet_leases",
			Help: "Total number of leases IP addresses of virtual network",
		},
		[]string{"vnet"},
	)
)

func RegisterVnetMetrics(vnets *virtualnetwork.Pool) {
	vNetsLeases.Reset()
	vNetsTotal.Reset()
	for _, v := range vnets.VirtualNetworks {
		var totalIPs int
		for _, ar := range v.ARs {
			totalIPs += ar.Size
		}
		vNetsTotal.WithLabelValues(v.Name).Set(float64(totalIPs))
		vNetsLeases.WithLabelValues(v.Name).Set(float64(v.UsedLeases))
	}

	vNetsTotalOnce.Do(func() {
		prometheus.MustRegister(vNetsTotal)

	})
	vNetsLeasesOnce.Do(func() {
		prometheus.MustRegister(vNetsLeases)
	})

}
