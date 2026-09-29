package executors

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Kiisanz/k11-addon-sdk/addonapi"
)

type FileSaveExecutor struct{}

func (e *FileSaveExecutor) Execute(ctx context.Context, input map[string]interface{}, emitter addonapi.EventEmitter) (map[string]interface{}, error) {
	destPath, ok := input["path"].(string)
	if !ok || destPath == "" {
		return nil, fmt.Errorf("missing or invalid 'path' input")
	}

	content := input["content"]
	if content == nil {
		return nil, fmt.Errorf("missing 'content' input")
	}

	// create directory if not exists
	err := os.MkdirAll(filepath.Dir(destPath), 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create directory: %v", err)
	}

	// We might receive the content directly as string or as an object referencing a local file
	if contentStr, ok := content.(string); ok {
		// Just a heuristic: if it looks like a /tmp/ file path, let's copy it.
		// Otherwise, write it directly. (For real K11, we'd have a formal Asset type)
		if filepath.IsAbs(contentStr) && func() bool { _, err := os.Stat(contentStr); return err == nil }() {
			// Copy file
			srcFile, err := os.Open(contentStr)
			if err != nil {
				return nil, err
			}
			defer srcFile.Close()

			// If destPath is an existing directory, append the original filename
			if info, err := os.Stat(destPath); err == nil && info.IsDir() {
				destPath = filepath.Join(destPath, filepath.Base(contentStr))
			}

			dstFile, err := os.Create(destPath)
			if err != nil {
				return nil, err
			}
			defer dstFile.Close()

			_, err = io.Copy(dstFile, srcFile)
			if err != nil {
				return nil, err
			}
		} else {
			// Write literal content
			if info, err := os.Stat(destPath); err == nil && info.IsDir() {
				destPath = filepath.Join(destPath, "output.txt")
			}
			err = os.WriteFile(destPath, []byte(contentStr), 0644)
			if err != nil {
				return nil, err
			}
		}
	} else {
		// Just fmt it
		if info, err := os.Stat(destPath); err == nil && info.IsDir() {
			destPath = filepath.Join(destPath, "output.txt")
		}
		err = os.WriteFile(destPath, []byte(fmt.Sprintf("%v", content)), 0644)
		if err != nil {
			return nil, err
		}
	}

	msg := fmt.Sprintf("File saved to %s", destPath)
	emitter.Emit(addonapi.RunEvent{
		Type:    "addonapi.log",
		Message: &msg,
	})

	return map[string]interface{}{"path": destPath}, nil
}
