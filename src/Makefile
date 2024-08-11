.PHONY: app mocks test
app:
	go run main.go apex_network_api

rm-mocks:
	rm -rf ./testutils/mocks.*

gen-mocks:
	mockery --all --output=testutils/mocks --case=underscore --keeptree

mocks: rm-mocks gen-mocks

test:
	go test -v -coverprofile=cover.out.tmp -coverpkg=./... ./...