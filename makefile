# cmd color presets
GREEN=\n\033[1;32;40m
RED=\n\033[1;31;40m
NC=\033[0m # No Color

PKG_LIST := $(shell go list ./... | grep -v config)

# run server
run:
	@/bin/sh -c "echo \"${GREEN}[SERVER START]${NC}\""
	go run main.go
.PHONY: run

# test
test:
	@/bin/sh -c "echo \"${GREEN}[TEST START]${NC}\""
	go test -short ${PKG_LIST}
.PHONY: test

# coverage
cov:
	@/bin/sh -c 'echo "${GREEN}[Calculate Test Coverages]${NC}"'
	@mkdir -p .public/coverage
	@gocov test ${PKG_LIST} | gocov-html > .public/coverage/index.html
	@gocov test ${PKG_LIST} | gocov-xml > .public/coverage/coverage.xml
	@gocov test ${PKG_LIST} | gocov report
.PHONY: coverage