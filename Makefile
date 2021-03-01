NAME := dde
MAINPKG := github.com/Lajule/dde
VERSION := 0.0.1

CP := cp
TARGETS := all run debug watch generate test bootstrap lint format clean clean-test

all:
	go build -ldflags="-s -X 'main.Version=$(VERSION)'" -tags "$(GOTAGS)" -o $(NAME) .

run:
	go run -tags "$(GOTAGS)" .

debug:
	dlv debug --build-flags "-tags '$(GOTAGS)'" $(MAINPKG)

watch:
	air -c .air.toml

generate:
	go generate

test:
	go test -tags "$(GOTAGS)" -v ./...

bootstrap:
	for subdir in $(SUBDIRS); do \
		for dir in $$(find $$subdir -type d); do \
			$(CP) Makefile.in $$dir/Makefile; \
		done \
	done

lint:
	golint ./...

format:
	for file in $$(find . -type f -name "*.go"); do \
		gofmt -s -w $$file; \
	done

clean: clean-test
	$(RM) $(NAME)

clean-test:
	go clean -testcache

.PHONY: $(TARGETS)
