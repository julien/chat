SHELL = /bin/bash

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))
app = $(current_dir)


all: clean $(app) run

$(app): 
	@go build

run:
	@echo ...run
	@PORT=3000 ./${app}

release:
	@echo ...release
	@go get

clean:
	@echo ...clean
	@rm -rf $(app)

test:
	@echo NO TESTS!

.PHONY: clean $(app) run release test

