ROOT_DIR=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

build:
	docker build -t ApiSetup ./

build-production:
	docker build -t ApiSetup ./ --build-arg app_env=production

serve:
	docker run -it -p 8080:8080 --net="host" -v ${ROOT_DIR}:/go/src/github.com/earlybird-SynchClient ApiSetup

test:
	docker run --entrypoint "run-tests.sh" ApiSetup
