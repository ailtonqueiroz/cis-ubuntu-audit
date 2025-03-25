.PHONY: build

build:
    @cd cmd && go build -o ../bin/cis-audit
