package connector

import (
	"fmt"
	"sync"

	"github.com/Haameed/nebula_exporter/exporter"

	"github.com/OpenNebula/one/src/oca/go/src/goca"
)

func GetVnet(client *goca.Controller, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		vnets, err := client.VirtualNetworks().Info()

		if err != nil {
			fmt.Println("Could not connect to opennebula check your inputs")
			fmt.Println(err.Error())
		}
		exporter.RegisterVnetMetrics(vnets)
	}()
}
