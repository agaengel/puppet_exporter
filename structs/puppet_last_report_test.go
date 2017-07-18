package structs

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestUnmarshall(t *testing.T) {
	assert := assert.New(t)
	testyamlPath := "../test/last_run_summary.yaml"
	data, err := ioutil.ReadFile(testyamlPath)

	if err != nil {
		t.Errorf("error reading testxml in Path: %v", testyamlPath)
	}

	parsedYaml, err := Unmarshall(data)
	if err != nil {
		t.Errorf("Parsing Error: %v", testyamlPath)
	}

	expect := &PuppetValues{
		Resources: Resources{
			Changed:          5,
			CorrectiveChange: 5,
			Failed:           0,
			FailedToRestart:  0,
			OutOfSync:        5,
			Restarted:        0,
			Scheduled:        0,
			Skipped:          0,
			Total:            335,
		},
	}

	assert.Equal(parsedYaml.Resources, expect.Resources)

	t.Log("Parsing of last_run_summary sucessfull")
}
