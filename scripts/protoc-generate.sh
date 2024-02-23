#!/usr/bin/env bash

# Root directory of app
ROOT_DIR=$(git rev-parse --show-toplevel)

PROTO_FILE_DIR=./gateway/backend-for-frontend/src/proto
API_PROTO_FILES=$(find ${PROTO_FILE_DIR} -type f -name '*.proto')

# User Service Generate
# echo "Generating code for User Service... "

OUT_DIR_USER_SEVICE="${ROOT_DIR}/micro-service/user-service/grpc"
PROTO_OUT_DIR_USER_SEVICE="./micro-service/user-service/grpc"

## Clean all existing generated files
rm -r "${OUT_DIR_USER_SEVICE}"
mkdir "${OUT_DIR_USER_SEVICE}"

## Generate code at User Service
protoc \
  -I=${PROTO_FILE_DIR} \
  --go_out=paths=source_relative:${PROTO_OUT_DIR_USER_SEVICE} \
  --go-grpc_out=paths=source_relative,require_unimplemented_servers=false:${PROTO_OUT_DIR_USER_SEVICE} \
  ${API_PROTO_FILES};

# echo "Generating code for User Service... Done"


# BFF Generate
# echo "Generating code for BFF..."
SRC_DIR="${ROOT_DIR}/gateway/backend-for-frontend/src/proto"

## Path to Protoc Plugin
PROTOC_GEN_TS_PATH="${ROOT_DIR}/gateway/backend-for-frontend/node_modules/.bin/protoc-gen-ts_proto"

## Directory to write generated code (.d.ts files)
OUT_DIR_BFF="${ROOT_DIR}/gateway/backend-for-frontend/src/grpc"

## Clean all existing generated files
rm -r "${OUT_DIR_BFF}"
mkdir "${OUT_DIR_BFF}"

## Generate code at BFF
# protoc \
#     --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
#     --ts_opt=esModuleInterop=true \
#     --ts_out="${OUT_DIR_BFF}" \
#     --grpc_out="${OUT_DIR_BFF}" \
#     --proto_path="${SRC_DIR}" \
#     $(find "${SRC_DIR}" -iname "*.proto")

protoc \
    --ts_proto_opt=nestJs=true \
    --plugin="${PROTOC_GEN_TS_PATH}" \
    --ts_proto_out="${OUT_DIR_BFF}" \
    --proto_path="${SRC_DIR}" \
     $(find "${SRC_DIR}" -iname "*.proto")

# echo "Generating code for BFF... Done"