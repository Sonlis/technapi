# -count 1 prevents go test from using cached tests results
#  trap is used to make sure the docker container is teared down
#  once the tests are done, whether or not they failed.
#  Some tests check if creating resources work, and it will fail
#  if the resource already exists.
.PHONY: test
test: compose
	@trap 'docker compose down && rm inventory.yaml zone-config.yaml' EXIT; \
	go test -count 1 ./... && python integration-test.py

.PHONY: compose
compose:
	docker compose up -d
