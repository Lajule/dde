VERSION := 0.0.0

TARGETS := all run debug watch test bootstrap lint format clean clean-test clean-watch

all:
	go build -ldflags="-s -X 'main.Version=$(VERSION)'" -tags "$(GOTAGS)" -o dde .

run:
	go run -tags "$(GOTAGS)" .

debug:
	dlv debug --build-flags "-tags '$(GOTAGS)'" github.com/Lajule/dde

watch:
	air -c .air.conf

test:
	go test -tags "$(GOTAGS)" -v ./...

bootstrap:
	for dir in $(SUBDIRS); do \
		$(RM) $$dir/Makefile; \
		{ \
			echo "TARGETS := $(TARGETS)"; \
			echo ""; \
			echo '$$(TARGETS):'; \
			echo '	$$(MAKE) -C .. $$@'; \
			echo ""; \
			echo '.PHONY = $$(TARGETS)'; \
		} >$$dir/Makefile; \
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

clean-watch:
	$(RM) .tmp

.PHONY: $(TARGETS)
