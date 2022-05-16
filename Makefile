setup:
	brew tap heroku/brew
	brew install go heroku

dev:
	tilt up

test-wip:
	WIP_ACCEPTANCE_ENABLED=true make test

test: unit-test integration-test acceptance-test

ut: unit-test
unit-test:
	@go test  -run 'Unit' ./src/...
	@echo '✅  unit tests successful'

it: integration-test
integration-test:
	@go test  -run 'Integration' ./src/...
	@echo '✅  integration tests successful'

at: acceptance-test
acceptance-test:
	@BASE_URL=http://localhost:8000 go test  ./acceptance/...
	@echo '✅  acceptance tests successful'

atd: acceptance-test-dev
acceptance-test-dev:
	@BASE_URL=https://boiling-forest-09153.herokuapp.com go test -v ./acceptance/...
	@echo '✅  acceptance tests successful'

deploy: test
	heroku container:login
	heroku container:push web -a boiling-forest-09153
	heroku container:release web -a boiling-forest-09153
	heroku ps:scale web=1
