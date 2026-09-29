package executors

import (
	"context"
	"fmt"
	"github.com/Kiisanz/k11-addon-sdk/pkg/addonapi"
	"time"
)

type DelayExecutor struct{}

func (e *DelayExecutor) Execute(ctx context.Context, input map[string]interface{}, emitter addonapi.EventEmitter) (map[string]interface{}, error) {
	durationMsRaw, ok := input["durationMs"]
	if !ok {
		return nil, fmt.Errorf("missing required input: durationMs")
	}

	// In Go, unmarshaled JSON numbers are usually float64
	durationFloat, ok := durationMsRaw.(float64)
	if !ok {
		return nil, fmt.Errorf("durationMs must be a number, got %T", durationMsRaw)
	}

	duration := time.Duration(durationFloat) * time.Millisecond

	select {
	case <-time.After(duration):
		// Done
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	return map[string]interface{}{}, nil
}
