.PHONY : build run test

PACKAGES = $(shell go list ./... | grep -v -e . | tr '\n' ',')
PATH_KEY = config/key/
PRIVATE_KEY = private.key
PUBLIC_KEY = public.pem

prepare:
	if ! [ -f .env ]; then cp env.example .env; fi;
	if ! [ -f $(PATH_KEY)$(PRIVATE_KEY) ]; then \
		./config/key/generate.sh 12345; \
		mv $(PRIVATE_KEY) $(PATH_KEY); \
		mv $(PUBLIC_KEY) $(PATH_KEY); \
	fi;

build:
	go build -o bin

run: prepare build
	./bin

docker: prepare
	docker build -t agungdpcms:latest .

test:
	if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@echo ">> running unit test and calculate coverage"
	@go test ./... -cover -coverprofile=coverage.txt -covermode=set -coverpkg=$(PACKAGES)
	@go tool cover -func=coverage.txt