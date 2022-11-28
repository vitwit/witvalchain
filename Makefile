#!/usr/bin/make -f

BUILDDIR ?= $(CURDIR)/build

install: go.sum
		export GOFLAGS='-buildmode=pie'
		export CGO_CPPFLAGS="-D_FORTIFY_SOURCE=2"
		export CGO_LDFLAGS="-Wl,-z,relro,-z,now -fstack-protector"
		go install $(BUILD_FLAGS) ./cmd/witvalchain

BUILD_TARGETS := build

build: BUILD_ARGS=-o $(BUILDDIR)/

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

containerProtoVer=v0.2
containerProtoImage=tendermintdev/sdk-proto-gen:$(containerProtoVer)
containerProtoGen=regen-ledger-proto-gen-$(containerProtoVer)
containerProtoFmt=regen-ledger-proto-fmt-$(containerProtoVer)
containerProtoGenSwagger=regen-ledger-proto-gen-swagger-$(containerProtoVer)

proto-all: proto-lint proto-format proto-gen proto-check-breaking

proto-gen:
	@echo "Generating Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(containerProtoImage) \
		sh ./scripts/protocgen.sh; fi