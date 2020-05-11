bin/nom: pkg/gen/ordersapi
	go build -o bin/nom

pkg/gen/ordersapi: orders.yaml bin/swagger
	rm -Rf $@
	mkdir -p $@
	bin/swagger generate client -q -f orders.yaml -t $@

# Instead of using a git submodule to get all of mymove for just one file, use
# curl to fetch the latest from github
orders.yaml:
	curl -sSo $@ https://raw.githubusercontent.com/transcom/mymove/master/swagger/orders.yaml

bin/swagger:
	go build -o bin/swagger github.com/go-swagger/go-swagger/cmd/swagger

# This target ensures that the pre-commit hook is installed and kept up to date
# if pre-commit updates.
.PHONY: ensure_pre_commit
ensure_pre_commit: .git/hooks/pre-commit ## Ensure pre-commit is installed
.git/hooks/pre-commit: /usr/local/bin/pre-commit
	pre-commit install
	pre-commit install-hooks

.PHONY: go_deps
go_deps: .go_deps.stamp ## Install Go dependencies
.go_deps.stamp: go.mod
	go mod tidy
	touch .go_deps.stamp

.PHONY: test
test: bin/nom
	go test ./...

.PHONY: clean
clean:
	rm -rf bin/
	rm -f orders.yaml
	rm -Rf pkg/gen/

.PHONY: default
default: bin/nom
