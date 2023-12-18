.PHONY: build run docker-build


build:
	go build -o ./bin/go-bank-analysis

run: build
	./bin/go-bank-analysis

docker-build:
	docker build -t go-bank-analysis:${VERSION} \
		-t go-bank-analysis:latest \
		-f dockerfile .

