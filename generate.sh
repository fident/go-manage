#!/bin/bash
# File: generate.sh
# Description: run to generate project dependencies

# Compile latest protocol buffers
echo -e "\033[4mCompiling protocol buffers...\033[0m"
command -v protoc >/dev/null 2>&1 || { echo "'protoc' tool is required but missing." >&2; exit 1; }
command -v protoc-gen-go >/dev/null 2>&1 || { echo "'protoc-gen-go' tool is required but missing." >&2; exit 1; }
protoc -I ./vendor/github.com/fident/proto --go_out=plugins=grpc:./fidentapi/ ./vendor/github.com/fident/proto/*.proto
(cd ./fidentapi && go fix ./...)