.PHONY: test
test: compose
	# -count 1 prevents go test from using cached tests results
	#  trap is used to make sure the docker container is teared down
	#  once the tests are done, whether or not they failed.
	#  Some tests check if creating resources work, and it will fail
	#  if the resource already exists.
	@trap 'docker compose down' EXIT; \
	go test -count 1 ./...

.PHONY: compose
compose:
	docker compose up -d
