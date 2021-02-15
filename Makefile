NAME := dde
MAINPKG := github.com/Lajule/dde
VERSION := 0.0.0

TARGETS := all run debug watch test bootstrap lint format clean clean-test

all:
	go build -ldflags="-s -X 'main.Version=$(VERSION)'" -tags "$(GOTAGS)" -o dde .

run:
	go run -tags "$(GOTAGS)" .

debug:
	dlv debug --build-flags "-tags '$(GOTAGS)'" $(MAINPKG)

watch:
	air -c .air.toml

test:
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

clean: clean-test clean-watch
	$(RM) dde

clean-test:
	go clean -testcache

.PHONY: $(TARGETS)
