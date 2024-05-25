package connector

import (
	"fmt"
	"sync"

	"github.com/Haameed/nebula_exporter/exporter"

	"github.com/OpenNebula/one/src/oca/go/src/goca"
)

func GetFreeCPU(client *goca.Controller, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		cpus, err := client.Hosts().Monitoring(0)
		if err != nil {
			fmt.Println(err)
		}
		exporter.RegisterFreeCPUMetrics(cpus, client)
	}()
}
