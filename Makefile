setup:
	brew install go

test-wip:
	WIP_ACCEPTANCE_ENABLED=true make test

test: unit-test integration-test acceptance-test

unit-test:
	go test -run 'Unit' ./src/...

integration-test:
	go test -run 'Integration' ./src/...

acceptance-test:
	go test ./acceptance/...
