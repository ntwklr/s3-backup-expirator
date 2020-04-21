VERSION=1.3.0

all: fmt combined

combined:
	go install .

release: tag release-deps 
	gox -ldflags "-X main.version=${VERSION}" -output="build/{{.Dir}}_{{.OS}}_{{.Arch}}" .

fmt:
	go fmt ./...

release-deps:
	go get github.com/mitchellh/gox

tag:
	git tag -a -m 'v${VERSION}' v${VERSION} && git push origin v${VERSION}