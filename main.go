// Copyright 2015 Oliver Fesseler
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// puppet exporter, exports metrics from puppet last_run_report.yaml
package main

import (
	"flag"
	"net/http"

	"fmt"
	"io/ioutil"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"

	"github.com/agaengel/puppet_exporter/structs"
)

const (
	namespace = "puppet"
)

var (
	versionConfig             = prometheus.NewDesc(prometheus.BuildFQName(namespace, "version", "config"), "Version config", []string{"puppet_version"}, nil)
	resourcesChanged          = prometheus.NewDesc(prometheus.BuildFQName(namespace, "resources", "changed"), "Resources changed", []string{"puppet_version"}, nil)
	resourcesCorrectiveChange = prometheus.NewDesc(prometheus.BuildFQName(namespace, "resources", "corrective_change"), "Resources corrective_change", []string{"puppet_version"}, nil)
	resourcesFailed           = prometheus.NewDesc(prometheus.BuildFQName(namespace, "resources", "failed"), "Resources failed", []string{"puppet_version"}, nil)
	resourcesFailedToRestart  = prometheus.NewDesc(prometheus.BuildFQName(namespace, "resources", "failed_to_restart"), "Resources failed_to_restart", []string{"puppet_version"}, nil)
	resourcesOutOfSync        = prometheus.NewDesc(prometheus.BuildFQName(namespace, "resources", "out_of_sync"), "Resources out_of_sync", []string{"puppet_version"}, nil)
	resourcesRestarted        = prometheus.NewDesc(prometheus.BuildFQName(namespace, "resources", "restarted"), "Resources restarted", []string{"puppet_version"}, nil)
	resourcesScheduled        = prometheus.NewDesc(prometheus.BuildFQName(namespace, "resources", "scheduled"), "Resources scheduled", []string{"puppet_version"}, nil)
	resourcesSkipped          = prometheus.NewDesc(prometheus.BuildFQName(namespace, "resources", "skipped"), "Resources skipped", []string{"puppet_version"}, nil)
	resourcesTotal            = prometheus.NewDesc(prometheus.BuildFQName(namespace, "resources", "total"), "Resources total", []string{"puppet_version"}, nil)
	changesTotal              = prometheus.NewDesc(prometheus.BuildFQName(namespace, "changes", "total"), "Changes total", []string{"puppet_version"}, nil)
	eventsFailure             = prometheus.NewDesc(prometheus.BuildFQName(namespace, "events", "failure"), "Events failure", []string{"puppet_version"}, nil)
	eventsSuccess             = prometheus.NewDesc(prometheus.BuildFQName(namespace, "events", "success"), "Events success", []string{"puppet_version"}, nil)
	eventsTotal               = prometheus.NewDesc(prometheus.BuildFQName(namespace, "events", "total"), "Events total", []string{"puppet_version"}, nil)

	times = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "times"),
		"Duration of the different resources",
		[]string{"puppet_version", "resource"}, nil,
	)
)

type exporter struct {
	puppetLastReportPath string
}

// Describe all the metrics exported by Puppet exporter. It implements prometheus.Collector.
func (v *exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- versionConfig
	ch <- resourcesChanged
	ch <- resourcesCorrectiveChange
	ch <- resourcesFailed
	ch <- resourcesFailedToRestart
	ch <- resourcesOutOfSync
	ch <- resourcesRestarted
	ch <- resourcesScheduled
	ch <- resourcesSkipped
	ch <- resourcesTotal
	ch <- changesTotal
	ch <- eventsFailure
	ch <- eventsSuccess
	ch <- eventsTotal
	ch <- times
}

// Collect collects all the metrics
func (v *exporter) Collect(ch chan<- prometheus.Metric) {

	// Collect metrics from volume info
	yamlFile, err := ioutil.ReadFile(v.puppetLastReportPath)
	if err != nil {
		log.Infof("yamlFile.Get err   #%v ", err)
	}

	puppetValues, err := structs.Unmarshall(yamlFile)

	ch <- prometheus.MustNewConstMetric(
		versionConfig,
		prometheus.GaugeValue,
		float64(puppetValues.Version.Config),
		puppetValues.Version.Puppet,
	)
	ch <- prometheus.MustNewConstMetric(
		resourcesChanged,
		prometheus.GaugeValue,
		float64(puppetValues.Resources.Changed),
		puppetValues.Version.Puppet,
	)
	ch <- prometheus.MustNewConstMetric(
		resourcesCorrectiveChange,
		prometheus.GaugeValue,
		float64(puppetValues.Resources.CorrectiveChange),
		puppetValues.Version.Puppet,
	)
	ch <- prometheus.MustNewConstMetric(
		resourcesFailed, prometheus.GaugeValue,
		float64(puppetValues.Resources.Failed),
		puppetValues.Version.Puppet,
	)
	ch <- prometheus.MustNewConstMetric(
		resourcesFailedToRestart,
		prometheus.GaugeValue,
		float64(puppetValues.Resources.FailedToRestart),
		puppetValues.Version.Puppet,
	)
	ch <- prometheus.MustNewConstMetric(
		resourcesOutOfSync,
		prometheus.GaugeValue,
		float64(puppetValues.Resources.OutOfSync),
		puppetValues.Version.Puppet,
	)
	ch <- prometheus.MustNewConstMetric(
		resourcesRestarted,
		prometheus.GaugeValue,
		float64(puppetValues.Resources.Restarted),
		puppetValues.Version.Puppet,
	)
	ch <- prometheus.MustNewConstMetric(
		resourcesScheduled,
		prometheus.GaugeValue,
		float64(puppetValues.Resources.Scheduled),
		puppetValues.Version.Puppet,
	)
	ch <- prometheus.MustNewConstMetric(
		resourcesSkipped,
		prometheus.GaugeValue,
		float64(puppetValues.Resources.Skipped),
		puppetValues.Version.Puppet,
	)
	ch <- prometheus.MustNewConstMetric(
		resourcesTotal,
		prometheus.GaugeValue,
		float64(puppetValues.Resources.Total),
		puppetValues.Version.Puppet,
	)
	ch <- prometheus.MustNewConstMetric(
		changesTotal,
		prometheus.GaugeValue,
		float64(puppetValues.Changes.Total),
		puppetValues.Version.Puppet,
	)
	ch <- prometheus.MustNewConstMetric(
		eventsFailure,
		prometheus.GaugeValue,
		float64(puppetValues.Events.Failure),
		puppetValues.Version.Puppet,
	)
	ch <- prometheus.MustNewConstMetric(
		eventsSuccess,
		prometheus.GaugeValue,
		float64(puppetValues.Events.Success),
		puppetValues.Version.Puppet,
	)
	ch <- prometheus.MustNewConstMetric(
		eventsTotal,
		prometheus.GaugeValue,
		float64(puppetValues.Events.Total),
		puppetValues.Version.Puppet,
	)

	for resource, time := range puppetValues.Time {
		ch <- prometheus.MustNewConstMetric(
			times,
			prometheus.GaugeValue,
			time,
			puppetValues.Version.Puppet,
			resource,
		)
	}
}

//newExporter initialises exporter
func newExporter(puppetLastReportPath string) (*exporter, error) {
	if len(puppetLastReportPath) < 1 {
		log.Fatalf("Puppet last Report path is wrong: %v", puppetLastReportPath)
	}
	_, err := os.Open(puppetLastReportPath) // For read access.
	if err != nil {
		log.Fatalf("Unable able to open file %v -- Message from os.Open: %v", puppetLastReportPath, err)
	}
	return &exporter{
		puppetLastReportPath: puppetLastReportPath,
	}, nil
}

func init() {
	prometheus.MustRegister(version.NewCollector("puppet_exporter"))
}

func versionInfo() {
	fmt.Println(version.Print("puppet_exporter"))
	os.Exit(0)
}

func main() {
	// commandline arguments
	var (
		metricPath           = flag.String("metrics-path", "/metrics", "URL Endpoint for metrics")
		puppetLastReportPath = flag.String("puppetLastReportPath", "/var/opt/puppetlabs/puppet/cache/state/last_run_summary.yaml", "Path to the last_run_summary.yaml")
		showVersion          = flag.Bool("version", false, "Prints version information")
		listenAddress        = flag.String("listen-address", ":9199", "The address to listen on for HTTP requests.")
	)
	flag.Parse()

	if *showVersion {
		versionInfo()
	}

	exporter, err := newExporter(*puppetLastReportPath)
	if err != nil {
		log.Errorf("Creating new Exporter went wrong, ... \n%v", err)
	}
	prometheus.MustRegister(exporter)

	log.Infoln("Puppet Metrics Exporter ", version.Info())
	log.Infoln("Build context", version.BuildContext())
	log.Info("metricPath=", *metricPath)

	http.Handle(*metricPath, promhttp.Handler())
	if *metricPath != "/" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Location", *metricPath)
			w.WriteHeader(301)
			w.Write([]byte(`<html>
			<head><title>Puppet Exporter v` + version.Version + `</title></head>
			<body>
			<h1>Puppet Exporter v` + version.Version + `</h1>
			<p><a href='` + *metricPath + `'>Metrics</a></p>
			</body>
			</html>
		`))
		})
	}
	log.Fatal(http.ListenAndServe(*listenAddress, nil))

}
