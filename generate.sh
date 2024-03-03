#!/bin/bash

export GO_POST_PROCESS_FILE="/usr/local/bin/gofmt -w"

OPEN_API_FILE="openapi/doc.json"

openapi-generator-cli validate -i "${OPEN_API_FILE}" && \
openapi-generator-cli generate \
    -i "${OPEN_API_FILE}" \
    -g go \
    -p packageName=ezpin \
    --git-host github.com \
    --git-repo-id ezpin-go \
    --git-user-id pourcheragh
