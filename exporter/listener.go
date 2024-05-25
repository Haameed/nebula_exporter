package exporter

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Listener(listenOn string, listenPort int) {
	prometheus.Unregister(collectors.NewGoCollector())
	http.Handle("/metrics", promhttp.Handler())
	port := fmt.Sprintf("%v:%v", listenOn, listenPort)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
