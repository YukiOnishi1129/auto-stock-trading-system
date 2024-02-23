#!/usr/bin/env bash

# Root directory of app
ROOT_DIR=$(git rev-parse --show-toplevel)

# Path to Protoc Plugin
PROTOC_GEN_TS_PATH="${ROOT_DIR}/gateway/backend-for-frontend/node_modules/.bin/protoc-gen-ts"

# Directory holding all .proto files
SRC_DIR="${ROOT_DIR}/pkg/proto"

# Directory to write generated code (.d.ts files)
OUT_DIR="${ROOT_DIR}/gateway/backend-for-frontend/pkg/grpc"

# Clean all existing generated files
rm -r "${OUT_DIR}"
mkdir "${OUT_DIR}"

# Generate all messages
protoc \
    --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
    --ts_opt=esModuleInterop=true \
    --ts_out="${OUT_DIR}" \
    --proto_path="${SRC_DIR}" \
    $(find "${SRC_DIR}" -iname "*.proto")