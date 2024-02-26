# Go compiler
GO := go

# Binary file output directory
OUT_DIR := bin

# Server binary file name
SERVER_BIN := $(OUT_DIR)/server

# Client binary file name
CLIENT_BIN := $(OUT_DIR)/client

# Path to the generate_protobuf.sh script
PROTOC_SCRIPT := scripts/generate_protobuf.sh
GRPC_GENERATED_DIR := internal/quiz/generated

.PHONY: all
all: generate server client

.PHONY: generate
generate:
	@$(PROTOC_SCRIPT)

.PHONY: server
server: generate $(SERVER_BIN)

.PHONY: client
client: generate $(CLIENT_BIN)

# Compile server
$(SERVER_BIN): apps/server/*.go internal/quiz/**/*.go
	@mkdir -p $(OUT_DIR)
	$(GO) build -o $(SERVER_BIN) ./apps/server

# Compile client
$(CLIENT_BIN): apps/client/*.go internal/quiz/**/*.go
	@mkdir -p $(OUT_DIR)
	$(GO) build -o $(CLIENT_BIN) ./apps/client

.PHONY: clean
clean:
	rm -rf $(OUT_DIR)
	rm -rf $(GRPC_GENERATED_DIR)

.PHONY: run-server
run-server: server
	$(SERVER_BIN)

.PHONY: run-client
run-client: client
	$(CLIENT_BIN)

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all        : Generate protobuf and build server and client"
	@echo "  generate   : Generate protobuf"
	@echo "  server     : Build server binary (depends on generate)"
	@echo "  client     : Build client binary"
	@echo "  clean      : Remove build artifacts"
	@echo "  run-server : Build and run the server"
	@echo "  run-client : Build and run the client"
