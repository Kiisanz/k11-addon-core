package executors

import (
	"context"
	"fmt"

	"github.com/Kiisanz/k11-addon-sdk/addonapi"
)

type DatasetAppendExecutor struct{}

type DatasetAppendManyExecutor struct{}

func (e *DatasetAppendExecutor) Execute(_ context.Context, input map[string]interface{}, emitter addonapi.EventEmitter) (map[string]interface{}, error) {
	datasetID, ok := input["datasetId"].(string)
	if !ok || datasetID == "" {
		return nil, fmt.Errorf("datasetId is required")
	}
	record := addonapi.UnwrapTypedValue(input["record"])
	recordMap, ok := record.(map[string]interface{})
	if !ok || recordMap == nil {
		return nil, fmt.Errorf("record must be an object or typed record")
	}
	emitter.Emit(addonapi.RunEvent{
		Type:      "dataset.record",
		DatasetID: datasetID,
		Data:      recordMap,
	})
	return map[string]interface{}{"record": addonapi.NewTypedValue("record", recordMap)}, nil
}

func (e *DatasetAppendManyExecutor) Execute(_ context.Context, input map[string]interface{}, emitter addonapi.EventEmitter) (map[string]interface{}, error) {
	datasetID, ok := input["datasetId"].(string)
	if !ok || datasetID == "" {
		return nil, fmt.Errorf("datasetId is required")
	}
	records, ok := addonapi.UnwrapTypedValue(input["records"]).([]interface{})
	if !ok {
		return nil, fmt.Errorf("records must be a record[]")
	}
	for index, value := range records {
		record, ok := addonapi.UnwrapTypedValue(value).(map[string]interface{})
		if !ok || record == nil {
			return nil, fmt.Errorf("records[%d] must be a record", index)
		}
		emitter.Emit(addonapi.RunEvent{Type: "dataset.record", DatasetID: datasetID, Data: record})
	}
	return map[string]interface{}{"count": addonapi.NewTypedValue("number", len(records))}, nil
}
