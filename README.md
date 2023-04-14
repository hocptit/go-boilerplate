# GZ locker API

## Run project

1. Run dev
   ###  With ubuntu
   - Install go : version > golang:1.19.3
   - `cp .env.example .env`
   - `make golangci-lint`
   - `make air`
   - `go install`
   - `make run-dev`
   ###  With windows
   - copy .env.example to .env
   - `install golangci-lint`
   - `go install github.com/cosmtrek/air@latest`
   - `go install`
   - `air -c .air-windows.toml`
   ###  With windows
2. Run product
   - Updating...

## Common rule
1. Do not code in the src/share folder
2. Run ```make lint``` before update code or merge/rebase code.
   - If lint return error `src/share/constant/base_constants.go:1: File is not `gofmt`-ed with `-s` (gofmt)`, run:
     - gofmt -s -w *

## Rule
1. Code using logger in API should use:
   ```go
   package test
   logger := getLogger.GetLogger().Logging
   logger.Info(utils.TID(c), "Content log")
   ```
2. Using panic for throw http exception:
   ```go
   package test
   panic(exception.BadRequestError(errorCode.InvalidParams, parseError(err)))
   ```
3. Throw error must include errorCode
4. Service not call service
5. Update later