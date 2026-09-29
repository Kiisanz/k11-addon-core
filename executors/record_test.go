package executors

import (
	"github.com/Kiisanz/k11-addon-sdk/addonapi"
	"testing"
)

func TestRecordBuilderBuildsTypedRecord(t *testing.T) {
	executor := &RecordBuilderExecutor{}
	outputs, err := executor.Execute(nil, map[string]interface{}{"fields": map[string]interface{}{
		"name": addonapi.NewTypedValue("string", "Product A"), "price": 129000,
	}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	record, ok := outputs["record"].(map[string]interface{})
	if !ok {
		t.Fatalf("record output = %#v", outputs)
	}
	if record["$type"] != "record" {
		t.Fatalf("record type = %#v", record["$type"])
	}
	value := record["value"].(map[string]interface{})
	if value["name"].(map[string]interface{})["value"] != "Product A" {
		t.Fatalf("record fields = %#v", value)
	}
}

func TestRecordBuilderRequiresFields(t *testing.T) {
	_, err := (&RecordBuilderExecutor{}).Execute(nil, map[string]interface{}{"fields": map[string]interface{}{}}, nil)
	if err == nil {
		t.Fatal("expected empty fields error")
	}
}
