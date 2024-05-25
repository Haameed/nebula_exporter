package exporter

import (
	"sync"

	"github.com/OpenNebula/one/src/oca/go/src/goca/schemas/vm"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	vmStatusOnce sync.Once
	vmCPUonce sync.Once
	vmMemoryOnce sync.Once
)

var (
	vmStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "opennebula_monitoring_vm_status",
			Help: "OpenNebula Host Status (0 = INIT, 1 = PENDING, 2 = HOLD, 3 = ACTIVE, 4 = STOPPED,  5 = SUSPENDED, 6 = DONE, 8 = POWEROFF, 9 = UNDEPLOYED, 10 = CLONING, 11 = CLONING_FAILURE)",
		},
		[]string{"hostname"},
	)
	vmCPU = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "opennebula_monitoring_vm_cpu",
			Help: "The number of CPUs allocated to VMs",
		},
		[]string{"hostname"},
	)
	vmMemory = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "opennebula_monitoring_vm_memory",
			Help: "The amount of memory allocated to VMs",
		},
		[]string{"hostname"},
	)
)

func RegisterVMMetrics(vms *vm.Pool) {
	vmStatus.Reset()
	vmCPU.Reset()
	vmMemory.Reset()
	for _, v := range vms.VMs {
		cpu, _ := v.Template.GetCPU()
		memory, _ := v.Template.GetMemory()
		vmStatus.WithLabelValues(v.Name).Set(float64(v.StateRaw))
		vmCPU.WithLabelValues(v.Name).Set(float64(cpu))
		vmMemory.WithLabelValues(v.Name).Set(float64(memory))
	}
	vmStatusOnce.Do(func() {
		prometheus.MustRegister(vmStatus)
	})
	vmCPUonce.Do(func() {
		prometheus.MustRegister(vmCPU)
	})
	vmMemoryOnce.Do(func() {
		prometheus.MustRegister(vmMemory)
	})
}
