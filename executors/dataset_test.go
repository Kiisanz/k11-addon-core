package executors

import (
	"context"
	"testing"

	"github.com/Kiisanz/k11-addon-sdk/pkg/addonapi"
)

type datasetEventCollector struct{ events []addonapi.RunEvent }

func (c *datasetEventCollector) Emit(event addonapi.RunEvent) { c.events = append(c.events, event) }

func TestDatasetAppendExecutorEmitsRecordEvent(t *testing.T) {
	collector := &datasetEventCollector{}
	executor := &DatasetAppendExecutor{}
	outputs, err := executor.Execute(nil, map[string]interface{}{
		"datasetId": "dataset-1",
		"record":    addonapi.NewTypedValue("record", map[string]interface{}{"title": "Product A"}),
	}, collector)
	if err != nil {
		t.Fatal(err)
	}
	if len(collector.events) != 1 || collector.events[0].Type != "dataset.record" {
		t.Fatalf("unexpected events: %+v", collector.events)
	}
	if collector.events[0].DatasetID != "dataset-1" || collector.events[0].Data["title"] != "Product A" {
		t.Fatalf("unexpected dataset event: %+v", collector.events[0])
	}
	if outputs["record"].(map[string]interface{})["$type"] != "record" {
		t.Fatalf("expected typed record output: %#v", outputs)
	}
}

func TestDatasetAppendManyExecutorEmitsEveryRecord(t *testing.T) {
	collector := &datasetEventCollector{}
	executor := &DatasetAppendManyExecutor{}
	outputs, err := executor.Execute(context.Background(), map[string]interface{}{
		"datasetId": "products",
		"records": addonapi.NewTypedValue("record[]", []interface{}{
			addonapi.NewTypedValue("record", map[string]interface{}{"name": "A"}),
			addonapi.NewTypedValue("record", map[string]interface{}{"name": "B"}),
		}),
	}, collector)
	if err != nil {
		t.Fatalf("append many failed: %v", err)
	}
	if len(collector.events) != 2 || len(collector.events[0].Data) == 0 || len(collector.events[1].Data) == 0 {
		t.Fatalf("events = %#v", collector.events)
	}
	if addonapi.UnwrapTypedValue(outputs["count"]) != 2 {
		t.Fatalf("outputs = %#v", outputs)
	}
}
