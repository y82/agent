/**
 * Copyright (c) F5, Inc.
 *
 * This source code is licensed under the Apache License, Version 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package metric

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

type MetricsProducer struct {
	dataChannel chan metricdata.Metrics
	metrics     []metricdata.Metrics
	metricsLock sync.RWMutex
}

func NewMetricsProducer() *MetricsProducer {
	return &MetricsProducer{
		dataChannel: make(chan metricdata.Metrics),
		metrics:     []metricdata.Metrics{},
		metricsLock: sync.RWMutex{},
	}
}

// Starts listening to metrics on its internal metrics channel.
func (hp *MetricsProducer) StartListen(ctx context.Context) {
	for {
		select {
		case metrics := <-hp.dataChannel:
			hp.metricsLock.Lock()
			hp.metrics = append(hp.metrics, metrics)
			hp.metricsLock.Unlock()
		case <-ctx.Done():
			return
		}
	}
}

// Pushes metrics to the MetricsProducer's internal channel.
func (hp *MetricsProducer) RecordMetrics(metrics metricdata.Metrics) {
	hp.dataChannel <- metrics
}

// Dumps the latest recorded metrics and reinitializes the metrics array.
func (hp *MetricsProducer) Produce(context.Context) ([]metricdata.ScopeMetrics, error) {
	hp.metricsLock.Lock()
	defer hp.metricsLock.Unlock()

	scopeMetrics := []metricdata.ScopeMetrics{
		{
			Scope: instrumentation.Scope{
				Name:    "github.com/agent/v3",
				Version: "v0.1",
			},
			Metrics: hp.metrics,
		},
	}

	hp.metrics = []metricdata.Metrics{}

	return scopeMetrics, nil
}
