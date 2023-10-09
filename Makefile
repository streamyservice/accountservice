wire_generate:
	 wire gen accountservice/config

clean_modcache:
	go clean -modcache

test:
	go test ./... -v

.PHONY: clean_modcache wire_generate test
