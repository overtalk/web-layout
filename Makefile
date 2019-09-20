TAG := $(shell git describe --exact-match --tags 2>/dev/null)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git rev-parse --short HEAD)

LDFLAGS := $(LDFLAGS) -s -w -X main.commit=$(COMMIT) -X main.branch=$(BRANCH) -X main.tag=$(TAG)

.PHONY: gen-model
gen-model:
	cd ./tools/generate-mysql-model/ && \
	sh ./gen-model.sh

.PHONY: example-server
example-server:
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS) -X main.name=Example_Serer" -o bin/example_server ./server/example/main.go