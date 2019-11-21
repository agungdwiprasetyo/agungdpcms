.PHONY : build docker deploy run test

GCP_PROJECT_ID = mantab-tenanan-le
APP_NAME = agungdpcms
ts = $(shell date +%Y%m%d%H%M%S)
IMAGE_TAG = $(ts)

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

docker: prepare build
	docker build -t $(APP_NAME):latest .

deploy:
	docker build -t $(APP_NAME):$(IMAGE_TAG) .
	docker tag $(APP_NAME):$(IMAGE_TAG) gcr.io/$(GCP_PROJECT_ID)/$(APP_NAME):$(IMAGE_TAG)
	docker push gcr.io/$(GCP_PROJECT_ID)/$(APP_NAME):$(IMAGE_TAG)
	kubectl set image deployment/$(APP_NAME) $(APP_NAME)-sha256=gcr.io/$(GCP_PROJECT_ID)/$(APP_NAME):$(IMAGE_TAG)

test: build
	if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@echo ">> running unit test and calculate coverage"
	@go test ./... -cover -coverprofile=coverage.txt -covermode=set -coverpkg=$(PACKAGES)
	@go tool cover -func=coverage.txt