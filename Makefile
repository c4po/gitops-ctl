export GO111MODULE=on
os = $(shell uname -s | awk '{print tolower($0)}')

REVISION := `git rev-parse HEAD`
BRANCH := `git rev-parse --abbrev-ref HEAD`
LDFLAGS=-ldflags "-X=github.com/c4po/gitops-ctl/cmd.Branch=$(BRANCH) \
-X=github.com/c4po/gitops-ctl/cmd.Revision=$(REVISION)"

.PHONY: build
build:
	go build $(LDFLAGS)

.PHONY: build-release
build-release:
	mkdir -p bin
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/windows-gitops-ctl
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/darwin-gitops-ctl
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/linux-gitops-ctl

.PHONY: test
test:
	go test -v ./...

# .PHONY: publish
# publish:
# 	gsutil cp bin/darwin-terraformctl gs://etsy-terraformctl-prod/${BRANCH_NAME}/darwin-terraformctl
# 	gsutil cp bin/linux-terraformctl gs://etsy-terraformctl-prod/${BRANCH_NAME}/linux-terraformctl

.PHONY: fmt
fmt:
	go mod tidy
	gofmt -s -w .

.PHONY: checkfmt
checkfmt:
	test -z "$$(gofmt -l .)"
