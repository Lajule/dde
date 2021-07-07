NAME := dde
PACKAGE := github.com/Lajule/dde
VERSION := 0.0.1
TARGETS := all run debug watch generate tidy test bootstrap lint format dist clean clean-test

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

tidy:
	go mod tidy

test:
	go test -tags "$(GOTAGS)" -v ./...

bootstrap:
	echo "tmp_dir = \".tmp\"\n\n[build]\ncmd = \"make\"\nbin = \"$(NAME)\"\ninclude_ext = [\"go\"]\nexclude_dir = []\ninclude_dir = []\nexclude_file = []\n\n[misc]\nclean_on_exit = true" >.air.toml
	find . -mindepth 1 -type d -exec sh -c "echo \"$(TARGETS):\n\t\\\$$(MAKE) -C .. \\\$$@\n\n.PHONY = $(TARGETS)\" >{}/Makefile" \;

lint:
	golint ./...

format:
	find . -type f -name "*.go" -exec gofmt -s -w {} \;

dist:
	touch tarball.tar.gz
	tar -czf tarball.tar.gz --exclude=tarball.tar.gz .

clean: clean-test
	$(RM) $(NAME)

clean-test:
	go clean -testcache

.PHONY: $(TARGETS)
