setup:
	brew install go

test-wip:
	WIP_ACCEPTANCE_ENABLED=true make test

test: unit-test integration-test acceptance-test

unit-test:
	go test -v -run 'Unit' ./src/...

integration-test:
	go test -v -run 'Integration' ./src/...

acceptance-test:
	go test -v ./acceptance/...
