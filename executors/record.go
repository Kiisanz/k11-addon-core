package executors

import (
	"context"
	"fmt"

	"github.com/Kiisanz/k11-addon-sdk/addonapi"
)

// RecordBuilderExecutor converts resolved named fields into the generic record type.
// It deliberately has no Dataset knowledge; persistence is handled by dataset.append.
type RecordBuilderExecutor struct{}

func (e *RecordBuilderExecutor) Execute(_ context.Context, input map[string]interface{}, _ addonapi.EventEmitter) (map[string]interface{}, error) {
	fieldsValue := addonapi.UnwrapTypedValue(input["fields"])
	fields, ok := fieldsValue.(map[string]interface{})
	if !ok || fields == nil {
		return nil, fmt.Errorf("fields must be an object")
	}
	if len(fields) == 0 {
		return nil, fmt.Errorf("at least one record field is required")
	}
	for name := range fields {
		if name == "" {
			return nil, fmt.Errorf("record field name cannot be empty")
		}
	}
	return map[string]interface{}{"record": addonapi.NewTypedValue("record", fields)}, nil
}
