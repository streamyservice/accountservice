install_commitlint:
	npm install -g @commitlint/cli @commitlint/config-conventional

precommit_install:
	pre-commit install

install_precommit:
	brew install pre-commit

wire_generate:
	 wire gen accountservice/config
clean_modcache:
	go clean -modcache

.PHONY: install_commitlint precommit_install install_precommit clean_modcache
