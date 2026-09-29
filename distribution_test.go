package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"testing"

	"github.com/Kiisanz/k11-addon-sdk/pkg/distribution"
)

func TestDistributionManifestSchema(t *testing.T) {
	b, err := os.ReadFile("distribution.json")
	if err != nil {
		t.Skip("distribution.json not generated yet")
	}

	var m distribution.DistributionManifest
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("could not parse distribution.json: %v", err)
	}

	if err := m.Validate(); err != nil {
		t.Errorf("manifest validation failed: %v", err)
	}
}

func TestDistributionArtifact(t *testing.T) {
	if os.Getenv("VERIFY_ARTIFACTS") != "1" {
		t.Skip("artifact verification disabled")
	}

	b, err := os.ReadFile("distribution.json")
	if err != nil {
		t.Fatalf("could not read distribution.json: %v", err)
	}

	var m distribution.DistributionManifest
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("failed to unmarshal distribution.json: %v", err)
	}

	if len(m.Assets) == 0 {
		t.Fatal("no assets found")
	}

	tarball, err := os.ReadFile("dist/core-linux-amd64.tar.gz")
	if err != nil {
		t.Fatalf("could not read tarball: %v", err)
	}

	if m.Assets[0].Size != int64(len(tarball)) {
		t.Errorf("expected size %d, got %d", m.Assets[0].Size, len(tarball))
	}

	hash := sha256.Sum256(tarball)
	actualHash := hex.EncodeToString(hash[:])
	if m.Assets[0].SHA256 != actualHash {
		t.Errorf("expected SHA-256 %s, got %s", m.Assets[0].SHA256, actualHash)
	}
}
