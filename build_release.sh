#!/bin/bash
set -e

TAG="${1:-${GITHUB_REF_NAME:-}}"
if [ -z "$TAG" ]; then
    echo "Error: missing release tag" >&2
    exit 1
fi
VERSION="${TAG#v}"

mkdir -p dist
GOOS=linux GOARCH=amd64 go build -o dist/core ./cmd/k11-addon-core
cd dist
tar -czf core-linux-amd64.tar.gz core
SIZE=$(stat -c%s core-linux-amd64.tar.gz)
SHA=$(sha256sum core-linux-amd64.tar.gz | awk '{print $1}')
cd ..

cat << JSON_EOF > distribution.json
{
  "schemaVersion": 1,
  "id": "core",
  "version": "${VERSION}",
  "api": {
    "addonApi": "v1",
    "editorSchemaVersion": 1,
    "minAgentVersion": "0.1.0"
  },
  "assets": [
    {
      "os": "linux",
      "arch": "amd64",
      "url": "https://github.com/Kiisanz/k11-addon-core/releases/download/${TAG}/core-linux-amd64.tar.gz",
      "sha256": "${SHA}",
      "size": ${SIZE},
      "format": "tar.gz",
      "executable": "core"
    }
  ]
}
JSON_EOF
