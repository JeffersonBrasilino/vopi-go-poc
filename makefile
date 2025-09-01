
make-deps-vendor:
	# create vendor directory
	go mod vendor
	# clean module cache
	go clean -modcache

start-dev:
	make make-deps-vendor
	docker compose up

test/coverage-html:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	
test/coverage-terminal:
	go test -cover ./...