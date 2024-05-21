package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Haameed/nebula_exporter/configparser"
	"github.com/Haameed/nebula_exporter/connector"
	"github.com/Haameed/nebula_exporter/exporter"
	"github.com/Haameed/nebula_exporter/input"
)

func main() {
	filepath := input.GetInput()
	if filepath != "" {
		endpoint, username, password, port, ssl, listenOn, listenPort, interval := configparser.ParseConfig(filepath)
		ticker := time.NewTicker(time.Duration(interval) * time.Second)
		defer ticker.Stop()
		fmt.Printf("Starting Prometheus listener on: %v:%v\n", listenOn, listenPort )
		go exporter.Listener(listenOn, listenPort)
		fmt.Printf("Connecting to: %v\n", endpoint )
		for range ticker.C {
			conectstr := connector.NewConnector(username, password, endpoint, port, ssl)
			clinet := connector.Connect(conectstr)
			wg := sync.WaitGroup{}
			connector.GetHosts(clinet, &wg)
			connector.GetVMs(clinet, &wg)
			connector.GetFreeCPU(clinet, &wg)
			connector.GetVnet(clinet, &wg)
			connector.GetDataStores(clinet, &wg)
			wg.Wait()
		}

	}

}
