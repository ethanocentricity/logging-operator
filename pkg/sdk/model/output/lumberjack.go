package output

import (
	"github.com/banzaicloud/operator-tools/pkg/secret"
	"github.com/banzaicloud/logging-operator/pkg/sdk/model/types"
)
// +docName:"Lumberjack output plugin for Fluentd"
//This plugin has been designed to output logs to lumberjack logging.
//
// #### Example output configurations
// ```
// spec:
//  lumberjack:
//    compartment: ocid1.compartment.oc1..aaaaaaaamepd3qlzatqffto3i32udpnyjytm7sp4dea6rny7e7wwkf3evmjq
//    namespace: app
// ```
type _docLumberjack interface{}

// +kubebuilder:object:generate=true

type LumberjackOutput struct {
	// Lumberjack logging compartment
	Compartment string `json:"compartment"`
	// Lumberjack namespace
	Namespace string `json:"namespace"`
	// +docLink:"Buffer,./buffer.md"
	Buffer *Buffer `json:"buffer,omitempty"`
}

func (l *LumberjackOutput) ToDirective(secretLoader secret.SecretLoader, id string) (types.Directive, error) {

	pluginType := "lumberjack"
	pluginID := id + "_" + pluginType
	lj := &types.OutputPlugin{
	    PluginMeta: types.PluginMeta{
	        Type:      pluginType,
	        Directive: "match",
	        Tag:       "**",
	        Id:        pluginID,
	    },
	}
	if params, err := types.NewStructToStringMapper(secretLoader).StringsMap(l); err != nil {
	    return nil, err
	} else {
	    lj.Params = params
        }

	if l.Buffer != nil {
	    if buffer, err := l.Buffer.ToDirective(secretLoader, pluginID); err != nil {
	        return nil, err
	    } else {
	        lj.SubDirectives = append(lj.SubDirectives, buffer)
	    }
	}
	return lj, nil
}
