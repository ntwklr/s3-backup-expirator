#!/bin/bash

VERSION=1.0.0

gox -ldflags "-X main.version=${VERSION}" -output="build/{{.Dir}}_{{.OS}}_{{.Arch}}"