package exporter

import (
	"sync"

	"github.com/OpenNebula/one/src/oca/go/src/goca/schemas/host"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	hostStatusOnce      sync.Once
	cpuTotalOnce        sync.Once
	cpuMaxOnce          sync.Once
	cpuUsageOnce        sync.Once
	memoryTotalOnce     sync.Once
	memoryAllocatedOnce sync.Once
	memoryMaxonce       sync.Once
)
var (
	hostStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "opennebula_monitoring_host_status",
			Help: "OpenNebula Host Status (0 = INIT, 1 = MONITORING_MONITORED, 2 = MONITORED, 3 = ERROR, 4 = DISABLED)",
		},
		[]string{"hostname", "cluster"},
	)

	cpuTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "opennebula_monitoring_cpu_total",
			Help: "Total Cpu threads",
		},
		[]string{"hostname", "cluster"},
	)
	cpuMax = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "opennebula_monitoring_cpu_max",
			Help: "Maximum CPU threads that are set for host",
		},
		[]string{"hostname", "cluster"},
	)
	cpuUsed = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "opennebula_monitoring_cpu_used",
			Help: "Number Of allocated CPUs to VMs",
		},
		[]string{"hostname", "cluster"},
	)
	memoryTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "opennebula_monitoring_memory_total",
			Help: "Total Memory of Host MB",
		},
		[]string{"hostname", "cluster"},
	)
	memoryAllocated = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "opennebula_monitoring_memory_allocated",
			Help: "Allocated memory MB",
		},
		[]string{"hostname", "cluster"},
	)
	memoryMax = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "opennebula_monitoring_memory_max",
			Help: "Maximum memory which can be allocated MB",
		},
		[]string{"hostname", "cluster"},
	)
)

func RegisterHostMetrics(hosts *host.Pool) {
	hostStatus.Reset()
	cpuTotal.Reset()
	cpuMax.Reset()
	cpuUsed.Reset()
	memoryTotal.Reset()
	memoryAllocated.Reset()
	memoryMax.Reset()
	for _, h := range hosts.Hosts {
		hostStatus.WithLabelValues(h.Name, h.Cluster).Set(float64(h.StateRaw))
		cpuTotal.WithLabelValues(h.Name, h.Cluster).Set(float64(h.Share.TotalCPU))
		cpuMax.WithLabelValues(h.Name, h.Cluster).Set(float64(h.Share.MaxCPU))
		cpuUsed.WithLabelValues(h.Name, h.Cluster).Set(float64(h.Share.CPUUsage))
		memoryTotal.WithLabelValues(h.Name, h.Cluster).Set(float64(h.Share.TotalMem))
		memoryAllocated.WithLabelValues(h.Name, h.Cluster).Set(float64(h.Share.MemUsage))
		memoryMax.WithLabelValues(h.Name, h.Cluster).Set(float64(h.Share.MaxMem))

	}
	hostStatusOnce.Do(func() {
		prometheus.MustRegister(hostStatus)

	})
	cpuTotalOnce.Do(func() {
		prometheus.MustRegister(cpuTotal)
	})
	cpuMaxOnce.Do(func() {
		prometheus.MustRegister(cpuMax)

	})
	cpuUsageOnce.Do(func() {
		prometheus.MustRegister(cpuUsed)

	})
	memoryTotalOnce.Do(func() {
		prometheus.MustRegister(memoryTotal)

	})
	memoryAllocatedOnce.Do(func() {
		prometheus.MustRegister(memoryAllocated)

	})
	memoryMaxonce.Do(func() {
		prometheus.MustRegister(memoryMax)

	})

}
