
make-deps-vendor:
	# create vendor directory
	go mod vendor
	# clean module cache
	go clean -modcache

start-dev:
	make make-deps-vendor
	docker compose up
