package structs

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"time"
)

func TestUnmarshallReport(t *testing.T) {
	assert := assert.New(t)
	testyamlPath := "../test/last_run_report.yaml"
	data, err := ioutil.ReadFile(testyamlPath)

	if err != nil {
		t.Errorf("error reading testxml in Path: %v", testyamlPath)
	}

	parsedYaml, err := UnmarshallReport(data)
	if err != nil {
		t.Errorf("Parsing Error: %v", testyamlPath)
	}

	expectedTime, err := time.Parse(time.RFC3339Nano, "2020-04-27T09:19:10.748030443+00:00")

	assert.Equal(parsedYaml.Time, expectedTime)

	t.Log("Parsing of last_run_report successful")
}
