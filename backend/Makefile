.PHONY: build
.DEFAULT_GOAL := build
COVERAGE_FORMAT := func

build:
	CGO_ENABLED=0 go build -v -o ./dist/holiday ./cmd/holiday

test:
	mkdir -p reports
	go test -cover ./... -coverprofile ./reports/coverage.out -coverpkg ./...
	go tool cover -$(COVERAGE_FORMAT) ./reports/coverage.out

lint:
	golangci-lint run ./... --timeout 5m0s

clean:
	golangci-lint run ./... --fix

log:
	git-chglog --config docs/.chglog/config.yml --output docs/CHANGELOG.md --next-tag $(tag)

release:
	git flow release start ${tag}
	sed -i "" 's/const version = "${SEMVER_REGEX}"/const version = "${tag}"/' ./cmd/holiday/main.go
	git-chglog --config docs/.chglog/config.yml --output docs/CHANGELOG.md --next-tag $(tag)
	golangci-lint run ./... --timeout 5m0s
	go test ./internal/... -test.count 3
	git add .
	git commit -m "bumped the version number"
	git flow release finish ${tag}

init:
	go mod tidy
	go mod download