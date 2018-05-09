PKGS=./cli/... ./pkg/...
PKGST=./cli ./pkg

.PHONY: install
install:
	go install ./cli/dy-bot

.PHONY: fmt
fmt:
	gofmt -d -l -w $(PKGST)

.PHONY: test
test:
	go test -v -cover $(PKGS)

.PHONY: dep-install
dep-install:
	dep ensure

.PHONY: dep-update
	dep ensure -update