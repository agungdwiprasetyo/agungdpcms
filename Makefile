.PHONY : build run

build:
	go build -o bin

run: build
	./bin

docker:
	if ! [ -f .env ]; then cp env.example .env; fi;
	docker build -t agungdpcms:latest .