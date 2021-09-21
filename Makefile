NAME := dde
PACKAGE := github.com/Lajule/dde
VERSION := 0.0.1
TARGETS := all run debug watch generate tidy test vet lint format clean bootstrap dist

all:
	go build -ldflags="-s -X 'main.Version=$(VERSION)'" -tags "$(GOTAGS)" -o $(NAME) .

run:
	go run -tags "$(GOTAGS)" .

debug:
	dlv debug --build-flags "-tags '$(GOTAGS)'" $(PACKAGE)

watch:
	air -c .air.toml

generate:
	go generate ./...

tidy:
	go mod tidy

test:
	go test -tags "$(GOTAGS)" -v ./...

vet:
	go vet ./...

lint:
	golint ./...

format:
	go fmt ./...

clean:
	go clean -r -i -cache -testcache -modcache

bootstrap:
	find . -mindepth 1 -type d -exec sh -c "echo \"$(TARGETS):\n\t\\\$$(MAKE) -C .. \\\$$@\n\n.PHONY = $(TARGETS)\" >{}/Makefile" \;

dist:
	touch $(NAME).tar.gz && tar -czf $(NAME).tar.gz --exclude=$(NAME).tar.gz .

.PHONY: $(TARGETS)
