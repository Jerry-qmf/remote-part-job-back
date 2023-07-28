# vars and envs
GO := go
GOOS := linux
GOARCH := amd64
GO_ENVS = CGO_ENABLED=0 GOOS=${GOOS} GOPROXY=direct GOSUMDB=off GOARCH=${GOARCH}

BUILD_TIME = $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_SHA := $(shell git rev-parse --short HEAD 2>/dev/null)
#VERSION := $(shell git describe --tags --always --dirty --match=v* 2> /dev/null)
VERSION = 0.1.0

GO_LD_FLAGS = \
	-X main.BuildTime=${BUILD_TIME} \
	-X main.GitSHA=${GIT_SHA} \
	-X main.Version=${VERSION}

#GO_BUILD_FLAGS = -mod=vendor -v -ldflags='${GO_LD_FLAGS}'
GO_BUILD_FLAGS = -v -ldflags='${GO_LD_FLAGS}'

BINARIES = remote-part-job
VERSION_NAME = ${BINARIES}-${VERSION}
DIST := dist/${VERSION_NAME}
TAR = ${VERSION_NAME}.tar.gz
PKGS = $(shell ${GO} list ./... | tr '\n' ',')
EXEC := ${DIST}/bin/${BINARIES}

# targets
all: clean build pkg

build:
	env ${GO_ENVS} ${GO} build ${GO_BUILD_FLAGS} -o ${EXEC} .

clean:
	rm -fr ${TAR}
	rm -fr dist

pkg:
	chmod a+x ${EXEC}
	cp -f script/start.sh ${DIST}/bin && chmod a+x ${DIST}/bin/start.sh
	cp -f script/build_systemd.sh ${DIST}/bin && chmod a+x ${DIST}/bin/build_systemd.sh
	cp -f config-live.json ${DIST}/config.json
	tar -zcf ${TAR} -C dist ${VERSION_NAME}
	@echo "build success!"

fmt:
	goimports -w cmd pkg

lint:
	tests/lint.sh

mod-tidy:
	${GO} mod tidy

${OUTPUT_PATH}:
	@mkdir -p ${OUTPUT_PATH}

test: coverage
	${GO} test ${GO_BUILD_FLAGS} -v -race ./...

coverage:
	@mkdir -p coverage
