package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

func registerLinkStatusForAdapters(reg *prometheus.Registry, cfg *config) {
	for name, adapter := range cfg.adapters {

		statusGauge := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "moca_link_status",
			Help: "Status of MoCA link",
		}, func() float64 {
			status, err := getMOCALinkStatus(&adapter)
			if err != nil {
				return 0
			}
			return float64(status)
		})

		prometheus.WrapRegistererWith(prometheus.Labels{"moca_adapter_name": name, "moca_adapter_address": adapter.MocaAdapterAddress}, reg).MustRegister(statusGauge)
	}
}
