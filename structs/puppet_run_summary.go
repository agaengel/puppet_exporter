package structs

import (
	"github.com/prometheus/common/log"
	"gopkg.in/yaml.v2"
)

//PuppetValues yaml root
type PuppetValues struct {
	Version   Version            `yaml:"version"`
	Resources Resources          `yaml:"resources"`
	Time      map[string]float64 `yaml:time`
	Changes   Changes            `yaml:changes`
	Events    Events             `yaml:events`
}

//Version in yaml
type Version struct {
	Config int32  `yaml:"config"`
	Puppet string `yaml:"puppet"`
}

//Resources in yaml
type Resources struct {
	Changed          int32 `yaml:"changed"`
	CorrectiveChange int32 `yaml:"corrective_change"`
	Failed           int32 `yaml:"failed"`
	FailedToRestart  int32 `yaml:"failed_to_restart"`
	OutOfSync        int32 `yaml:"out_of_sync"`
	Restarted        int32 `yaml:"restarted"`
	Scheduled        int32 `yaml:"scheduled"`
	Skipped          int32 `yaml:"skipped"`
	Total            int32 `yaml:"total"`
}

//Changes in yaml
type Changes struct {
	Total int32 `yaml:"total"`
}

//Events in yaml
type Events struct {
	Failure int32 `yaml:"failure"`
	Success int32 `yaml:"success"`
	Total   int32 `yaml:"total"`
}

//UnmarshallSummary yamlfile
func UnmarshallSummary(yamlInput []byte) (PuppetValues, error) {
	puppetValues := PuppetValues{}

	err := yaml.Unmarshal(yamlInput, &puppetValues)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return puppetValues, nil
}
