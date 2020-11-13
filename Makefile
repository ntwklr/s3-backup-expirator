VERSION=1.5.0

all: fmt combined

combined:
	go install .

release: tag release-deps 
	gox -ldflags -osarch "!darwin/386" "-X main.version=${VERSION}" -output="build/{{.Dir}}_{{.OS}}_{{.Arch}}" .

fmt:
	go fmt ./...

release-deps:
	go get github.com/mitchellh/gox

tag:
	git tag -a -m 'v${VERSION}' v${VERSION}

push:
	git push origin v${VERSION}