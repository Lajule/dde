NAME := dde
PACKAGE := github.com/Lajule/dde
VERSION := 0.0.1
TARGETS := all run debug watch generate test bootstrap lint format tarball clean clean-test

all:
	go build -ldflags="-s -X 'main.Version=$(VERSION)'" -tags "$(GOTAGS)" -o $(NAME) .

run:
	go run -tags "$(GOTAGS)" .

debug:
	dlv debug --build-flags "-tags '$(GOTAGS)'" $(PACKAGE)

watch:
	air -c .air.toml

generate:
	go generate

test:
	go test -tags "$(GOTAGS)" -v ./...

bootstrap:
	echo 'tmp_dir = ".tmp"\n\n[build]\ncmd = "make"\nbin = "$(NAME)"\ninclude_ext = ["go"]\nexclude_dir = []\ninclude_dir = []\nexclude_file = []\n\n[misc]\nclean_on_exit = true' >.air.toml
	for subdir in $(SUBDIRS); do \
		for dir in $$(find $$subdir -type d); do \
			echo 'TARGETS := $(TARGETS)\n\n$$(TARGETS):\n\t$$(MAKE) -C .. $$@\n\n.PHONY = $$(TARGETS)' >$$dir/Makefile; \
		done \
	done

lint:
	golint ./...

format:
	for file in $$(find . -type f -name "*.go"); do \
		gofmt -s -w $$file; \
	done

tarball:
	touch tarball.tar.gz
	tar -czf tarball.tar.gz --exclude=tarball.tar.gz .

clean: clean-test
	$(RM) $(NAME)

clean-test:
	go clean -testcache

.PHONY: $(TARGETS)
