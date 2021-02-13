NAME := dde
MAINPKG := github.com/Lajule/dde
VERSION := 0.0.0

TARGETS := all run debug watch test bootstrap lint format clean clean-test

all:
	go build -ldflags="-s -X 'main.Version=$(VERSION)'" -tags "$(GOTAGS)" -o $(CURDIR)/dde $(CURDIR)

run:
	go run -tags "$(GOTAGS)" $(CURDIR)

debug:
	dlv debug --build-flags "-tags '$(GOTAGS)'" $(MAINPKG)

watch:
	air -c .air.conf

test:
	go test -tags "$(GOTAGS)" -v $(CURDIR)/...

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
	golint $(CURDIR)/...

format:
	for file in $$(find $(CURDIR) -type f -name "*.go"); do \
		gofmt -s -w $$file; \
	done

clean: clean-test clean-watch
	$(RM) $(CURDIR)/dde

clean-test:
	go clean -testcache

.PHONY: $(TARGETS)
