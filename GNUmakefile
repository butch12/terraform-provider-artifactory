TEST?=./...
PRODUCT=artifactory
GO_ARCH=$(shell go env GOARCH)
TARGET_ARCH=$(shell go env GOOS)_${GO_ARCH}
GORELEASER_ARCH=${TARGET_ARCH}

ifeq ($(GO_ARCH), amd64)
GORELEASER_ARCH=${TARGET_ARCH}_$(shell go env GOAMD64)
endif
PKG_NAME=pkg/artifactory
# if this path ever changes, you need to also update the 'ldflags' value in .goreleaser.yml
PKG_VERSION_PATH=github.com/jfrog/terraform-provider-${PRODUCT}/${PKG_NAME}
PROVIDER_VERSION?=$(shell git tag --sort=-creatordate | head -1 | sed  -n 's/v\([0-9]*\).\([0-9]*\).\([0-9]*\)/\1.\2.\3/p')
NEXT_PROVIDER_VERSION := $(shell echo ${PROVIDER_VERSION}| awk -F '.' '{print $$1 "." $$2 "." $$3 +1 }' )
BUILD_PATH=terraform.d/plugins/registry.terraform.io/jfrog/${PRODUCT}/${NEXT_PROVIDER_VERSION}/${TARGET_ARCH}
SONAR_SCANNER_VERSION?=4.7.0.2747
SONAR_SCANNER_HOME?=$HOME/.sonar/sonar-scanner-$SONAR_SCANNER_VERSION-macosx

default: build

install: clean build
	rm -fR .terraform.d && \
	mkdir -p ${BUILD_PATH} && \
		mv -v dist/terraform-provider-${PRODUCT}_${GORELEASER_ARCH}/terraform-provider-${PRODUCT}_v${NEXT_PROVIDER_VERSION}* ${BUILD_PATH} && \
		rm -f .terraform.lock.hcl && \
		sed -i.bak 's/version = ".*"/version = "${NEXT_PROVIDER_VERSION}"/' sample.tf && rm sample.tf.bak && \
		terraform init

clean:
	rm -fR dist terraform.d/ .terraform terraform.tfstate* terraform.d/ .terraform.lock.hcl

release:
	@git tag ${NEXT_PROVIDER_VERSION} && git push --mirror
	@echo "Pushed ${NEXT_PROVIDER_VERSION}"
	GOPROXY=https://proxy.golang.org GO111MODULE=on go get github.com/jfrog/terraform-provider-${PRODUCT}@v${NEXT_PROVIDER_VERSION}
	@echo "Updated pkg cache"

update_pkg_cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on go get github.com/jfrog/terraform-provider-${PRODUCT}@v${PROVIDER_VERSION}

build: fmt
	GORELEASER_CURRENT_TAG=${NEXT_PROVIDER_VERSION} goreleaser build --single-target --rm-dist --snapshot

test:
	@echo "==> Starting unit tests"
	go test $(TEST) -timeout=30s -parallel=4

attach:
	dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient attach $$(pgrep terraform-provider-${PRODUCT})

acceptance: fmt
	export TF_ACC=true && \
		go test -cover -coverprofile=coverage.txt -ldflags="-X '${PKG_VERSION_PATH}.Version=${NEXT_PROVIDER_VERSION}-test'" -v -p 1 -parallel 20 -timeout 20m ./pkg/...

coverage:
	go tool cover -html=coverage.txt

scan:
	${SONAR_SCANNER_HOME}/bin/sonar-scanner -Dsonar.projectVersion=${PROVIDER_VERSION} -Dsonar.go.coverage.reportPaths=coverage.txt

fmt:
	@echo "==> Fixing source code with gofmt..."
	@go fmt ./...

.PHONY: build fmt
