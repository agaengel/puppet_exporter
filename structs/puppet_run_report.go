package structs

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Report yaml root
type Report struct {
	Time time.Time `yaml:"time"`
}

// UnmarshallReport yamlfile
func UnmarshallReport(yamlInput []byte) (Report, error) {
	report := Report{}

	err := yaml.Unmarshal(yamlInput, &report)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return report, nil
}
