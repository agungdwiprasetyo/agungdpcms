.PHONY : build run test cover

TEST_PACKAGES = ./src/resume/delivery \
				./src/resume/repository

build:
	go build -o bin

run: build
	./bin

docker:
	if ! [ -f .env ]; then cp env.example .env; fi;
	docker build -t agungdpcms:latest .

test:
	$(foreach pkg, $(TEST_PACKAGES),\
	go test $(pkg);)

cover:
	if [ -f coverage.txt ]; then rm coverage.txt; fi;
	$(foreach pkg, $(TEST_PACKAGES), \
	go test -coverprofile=coverage.out -covermode=atomic $(pkg); \
	tail -n +2 coverage.out >> coverage.txt;)
	rm coverage.out