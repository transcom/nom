nom: pkg/gen/ordersapi
	go build

pkg/gen/ordersapi: orders.yaml bin/swagger
	rm -Rf $@
	mkdir -p $@
	bin/swagger generate client -f orders.yaml -t $@

# Instead of using a git submodule to get all of mymove for just one file, use
# curl to fetch the latest from github
orders.yaml:
	curl -o $@ https://raw.githubusercontent.com/transcom/mymove/master/swagger/orders.yaml

bin/swagger: vendor/github.com/go-swagger/go-swagger/cmd/swagger
	go build -i -o $@ ./$<

vendor/github.com/go-swagger/go-swagger/cmd/swagger:
	dep ensure -vendor-only

clean:
	rm -f bin/swagger
	rm -f orders.yaml
	rm -Rf vendor
	rm -Rf pkg/gen

.PHONY: nom clean
