.PHONY: build run docker-build


build:
	go build -o ./bin/go-bank-analysis

run: build
	./bin/go-bank-analysis

docker-build:
	docker build -t go-bank-analysis:${VERSION} \
		-t go-bank-analysis:latest \
		-f dockerfile .

nordi-login:
	curl -X POST http://localhost:8080/nordigen/login

list-bank:
	curl -X GET http://localhost:8080/nordigen/banks

agree: 
	curl -X POST http://localhost:8080/nordigen/agree \
	  -d "{\"institution_id\": \"LABANQUEPOSTALE_PSSTFRPP\",
  		  \"max_historical_days\": \"90\",
		  \"access_valid_for_days\": \"30\",
		  \"access_scope\": [\"balances\", \"details\", \"transactions\"] }" | jq


record-bank:
	curl -X POST http://localhost:8080/nordigen/record-bank \ 
		-d "{\"redirect\": \"http://www.yourwebpage.com\",
		  	\"institution_id\": \"LABANQUEPOSTALE_PSSTFRPP\",
		  	\"reference\": \"11111\",
		  	\"agreement\": \"045f1b2d-e7af-4a88-87fe-5fb7f6e29bc9\",
		  	\"user_language\":\"EN\"}"

transactions:
	curl -X GET  http://localhost:8080/nordigen/501c5b59-4537-44e8-9677-f9e48e2ff139/transactions | jq | more

balances:
	curl -X GET  http://localhost:8080/nordigen/501c5b59-4537-44e8-9677-f9e48e2ff139/balances

start-db:
	sudo docker run --name my-db -e POSTGRES_PASSWORD=monpassword \
		-e POSTGRES_DB=bank \
		-d postgres:16
