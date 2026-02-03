dep:
	go install google.golang.org/protobuf/cmd/protoc-gen-go

build:
	go build ./validator/...

nuke:
	(cd validator && make nuke)

regenerate-tests:
	(cd validator && make regenerate-tests)

regenerate-all:
	(cd validator && make regenerate-all)

regenerate-paper-benchmarks:
	(cd validator && make regenerate-paper-benchmarks)
