package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type adapter struct {
	MocaAdapterAddress string `toml:"AdapterAddress"`
	MocaUser           string `toml:"User"`
	MocaPass           string `toml:"Password"`
}

type config struct {
	adapters map[string]adapter
}

func readConfig(file string) (*config, error) {

	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	cfg := new(config)
	// v := map[string]interface{}{}
	err = toml.Unmarshal(b, &(cfg.adapters))
	if err != nil {
		return nil, err
	}

	// for name, adapterInfo := range v {

	// }

	return cfg, nil
}

func main() {
	// Expose the registered metrics via HTTP.
	var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
	var configFile = flag.String("config", "configs/moca.toml", "The config file (in toml format) describing the adapters to monitor")

	flag.Parse()

	if *configFile == "" {
		log.Fatal("Config file was empty")
	}

	cfg, err := readConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	reg := prometheus.NewRegistry()
	registerLinkStatusForAdapters(reg, cfg)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
