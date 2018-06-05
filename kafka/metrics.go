package main

import (
	"github.com/Shopify/sarama"
	"github.com/rcrowley/go-metrics"
	"os"
)

func main() {
	// Our application registry
	appMetricRegistry := metrics.NewRegistry()
	appGauge := metrics.GetOrRegisterGauge("m1", appMetricRegistry)
	appGauge.Update(1)

	config := sarama.NewConfig()
	// Use a prefix registry instead of the default local one
	config.MetricRegistry = metrics.NewPrefixedChildRegistry(appMetricRegistry, "sarama.")

	// Simulate a metric created by sarama without starting a broker
	saramaGauge := metrics.GetOrRegisterGauge("m2", config.MetricRegistry)
	saramaGauge.Update(2)

	metrics.WriteOnce(appMetricRegistry, os.Stdout)
}
