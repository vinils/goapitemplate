# Binary base name
APP_NAME=myapp

# Output directory
BUILD_DIR?=bin

OUTPUT?=$(BUILD_DIR)/$(APP_NAME)

# Platforms to build for linux/amd64, linux/arm64, windows/amd64, darwin/amd64
PLATFORM?=windows/amd64
GOOS=$(word 1,$(subst /, ,$(PLATFORM)))
GOARCH=$(word 2,$(subst /, ,$(PLATFORM)))



GO := go
GO_BUILD := $(GO) build
GO_RUN := $(GO) run
GO_TEST := $(GO) test
GO_CLEAN := $(GO) clean
GO_MOD := $(GO) mod
GO_FMT := $(GO) fmt

# Default target
.PHONY: all
all: build

# Build for current OS
.PHONY: build
build:
	@echo $(OUTPUT)
	$(GO_BUILD) -o $(OUTPUT) .

# # Build for current OS
# .PHONY: build-linux
# build:
# 	if [ "${OS}" = "windows" ]; then OUTPUT=$${OUTPUT}.exe; fi; \
# 	CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} $(GO_BUILD) -o ${OUTPUT} .


# Run the application
.PHONY: run
run:
	$(GO_RUN) .

# Test the application
.PHONY: test
test:
	$(GO_TEST) -v -coverprofile=coverage.txt -covermode count ./...

# Test the application
.PHONY: coverage
test:
	go tool cover -html coverage.txt

# Format the code
.PHONY: fmt
fmt:
	$(GO_FMT) ./...

# Lint the code (requires golangci-lint)
.PHONY: lint
lint:
	golangci-lint run

# Clean build files
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# Tidy modules
.PHONY: tidy
tidy:
	$(GO_MOD) tidy
