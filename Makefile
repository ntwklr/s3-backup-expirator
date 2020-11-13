VERSION=1.5.0

all: fmt combined

combined:
	go install .

release: tag release-deps release-build
	git push origin v${VERSION}

fmt:
	go fmt ./...

tag:
	git tag -a -m 'v${VERSION}' v${VERSION}

release-deps:
	go get github.com/mitchellh/gox

release-build:
	gox -ldflags "-X main.version=${VERSION}" -osarch "!darwin/386" -output="build/{{.Dir}}_{{.OS}}_{{.Arch}}" .