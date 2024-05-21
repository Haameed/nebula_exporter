package connector

import (
	"fmt"
	"sync"

	"github.com/Haameed/nebula_exporter/exporter"

	"github.com/OpenNebula/one/src/oca/go/src/goca"
)

func GetVMs(client *goca.Controller, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		vms, err := client.VMs().Info()
		if err != nil {
			fmt.Println("could not connect to opennebula check your inputs")
			fmt.Println(err.Error())
		}
		exporter.RegisterVMMetrics(vms)

	}()

}
