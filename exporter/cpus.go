package exporter

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/OpenNebula/one/src/oca/go/src/goca"
	"github.com/OpenNebula/one/src/oca/go/src/goca/dynamic"
	"github.com/OpenNebula/one/src/oca/go/src/goca/schemas/host"
	"github.com/prometheus/client_golang/prometheus"
)

var freeCPUOnce sync.Once

var freeCPU = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "opennebula_monitoring_host_free_cpu",
		Help: "Free CPU",
	},
	[]string{"hostname", "cluster"},
)

type cpumMetricInfo struct {
	free     int
	hostinfo *host.Host
}

func RegisterFreeCPUMetrics(cpus *host.PoolMonitoring, client *goca.Controller) {
	wg := sync.WaitGroup{}
	ch := make(chan *cpumMetricInfo, len(cpus.Records))
	for _, c := range cpus.Records {
		wg.Add(1)
		go func(c dynamic.Template) {
			defer wg.Done()
			id, err := c.GetInt("ID")
			if err != nil {
				fmt.Println(err.Error())
			}
			hostinfo, _ := client.Host(id).Info(true)
			free, _ := c.GetStrFromVec("CAPACITY", "FREE_CPU")
			freeInt, _ := strconv.Atoi(free)
			ch <- &cpumMetricInfo{
				free:     freeInt,
				hostinfo: hostinfo,
			}
		}(c)
	}
	wg.Wait()
	close(ch)
	freeCPU.Reset()
	for m := range ch {
		hostname := m.hostinfo.Name
		cluster := m.hostinfo.Cluster
		freeCPU.WithLabelValues(hostname, cluster).Set(float64(m.free))

	}
	freeCPUOnce.Do(func() {
		prometheus.MustRegister(freeCPU)
	})
}
