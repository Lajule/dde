NAME := dde
MAINPKG := github.com/Lajule/dde
VERSION := 0.0.0

TARGETS := all run debug watch generate test bootstrap lint format clean clean-test

all: generate
	go build -ldflags="-s -X 'main.Version=$(VERSION)'" -tags "$(GOTAGS)" -o dde .

run: generate
	go run -tags "$(GOTAGS)" .

debug: generate
	dlv debug --build-flags "-tags '$(GOTAGS)'" $(MAINPKG)

watch:
	air -c .air.toml

generate:
	go generate

test: generate
	go test -tags "$(GOTAGS)" -v ./...

bootstrap:
	for subdir in $(SUBDIRS); do \
		for dir in $$(find $$subdir -type d); do \
			$(RM) $$dir/Makefile; \
			{ \
				echo "TARGETS := $(TARGETS)"; \
				echo ""; \
				echo '$$(TARGETS):'; \
				echo '	$$(MAKE) -C .. $$@'; \
				echo ""; \
				echo '.PHONY = $$(TARGETS)'; \
			} >$$dir/Makefile; \
		done \
	done

lint:
	golint ./...

format:
	for file in $$(find . -type f -name "*.go"); do \
		gofmt -s -w $$file; \
	done

clean: clean-test
	$(RM) dde

clean-test:
	go clean -testcache

.PHONY: $(TARGETS)
