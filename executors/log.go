package executors

import (
	"context"
	"fmt"
	"github.com/Kiisanz/k11-addon-sdk/addonapi"
)

type LogExecutor struct{}

func (e *LogExecutor) Execute(ctx context.Context, input map[string]interface{}, emitter addonapi.EventEmitter) (map[string]interface{}, error) {
	message, ok := input["message"]
	if !ok {
		return nil, fmt.Errorf("missing required input: message")
	}

	msgStr := fmt.Sprintf("%v", message)
	levelStr := "info"

	// Emit the structured log event!
	emitter.Emit(addonapi.RunEvent{
		Type:    "addonapi.log",
		Message: &msgStr,
		Level:   &levelStr,
	})

	fmt.Printf("[WORKFLOW LOG] %v\n", message)

	return map[string]interface{}{}, nil
}
