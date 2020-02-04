package output_test

import (
	"testing"

	"github.com/banzaicloud/logging-operator/pkg/sdk/model/output"
	"github.com/banzaicloud/logging-operator/pkg/sdk/model/render"
	"github.com/ghodss/yaml"
)

func TestLumberjack(t *testing.T) {
	CONFIG := []byte(`
compartment: test..compartment
namespace: test.ns
`)
	expected := `
  <match **>
    @type lumberjack
    @id test_lumberjack
    compartment test..compartment
    namespace test.ns
  </match>
`
	l := &output.LumberjackOutput{}
	yaml.Unmarshal(CONFIG, l)
	test := render.NewOutputPluginTest(t, l)
	test.DiffResult(expected)
}