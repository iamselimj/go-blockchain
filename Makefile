NAME := chain
VERSION := 0.1.0

BINARY := bin/$(NAME)
COMPILER := go
PKG := ./...
GOFLAGS := -v -ldflags "-X 'main.version=$(VERSION)' -X 'main.name=$(NAME)'"

build:
	@echo "🔨 Building $(NAME)..."
	@$(COMPILER) build $(GOFLAGS) -o $(BINARY) .

run: build
	@echo "🚀 Running $(NAME)..."
	@./$(BINARY)

test:
	@echo "🧪 Running tests..."
	@$(COMPILER) test $(PKG)

lint:
	@command -v staticcheck >/dev/null 2>&1 && staticcheck $(PKG) || echo "ℹ️  staticcheck non installé"

fmt:
	@echo "🧹 Formatting code..."
	@$(COMPILER) fmt $(PKG)

tidy:
	@echo "🧼 Tidying go.mod/go.sum..."
	@$(COMPILER) mod tidy

clean:
	@echo "🗑️  Cleaning binary..."
	@rm -f $(BINARY)

fclean: clean
	@echo "💣 Full clean (binary, bin, cache)..."
	@rm -rf bin
	@$(COMPILER) clean -cache -modcache -testcache

re: fclean build

install: build
	@echo "📦 Installing $(NAME)..."
	@$(COMPILER) install $(GOFLAGS)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'

.PHONY: build run test lint fmt tidy clean fclean re install help