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

.PHONY: clean
clean:
	rm -rf bin/
	rm -f orders.yaml
	rm -Rf pkg/gen/

.PHONY: default
default: bin/nom
