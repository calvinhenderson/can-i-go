## Variables

MAIN_PACKAGE_PATH := ./cmd/app
BINARY_NAME := can-i-go

## Usage

.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'


## Helpers

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read resp && [ $${resp:-N} = y ]

.PHONY: no-dirty
no-dirty:
	git diff --exit-code


## Quality Control

.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...

.PHONY: lint
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2 run

.PHONY: check
check: tidy lint audit no-dirty

## Testing

.PHONY: test
test: build
	go test -v -race -buildvcs ./...

.PHONY: test/cover
test/cover: build
	go test -v -race -buildvcs ./...
	go tool cover -html=/tmp/coverage.out


## Development

.PHONY: deps
deps:
	go get ./...

.PHONY: build
build: deps
	go build -o=/tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

.PHONY: run
run: build
	/tmp/bin/${BINARY_NAME}

.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" --build.bin "/tmp/bin/${BINARY_NAME}" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"


# vim: set ts=2 sts=2 sw=2 noet
