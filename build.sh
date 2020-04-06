#!/bin/bash

VERSION=0.0.1

gox -ldflags "-X main.version=${VERSION}" -output="build/{{.Dir}}_{{.OS}}_{{.Arch}}"