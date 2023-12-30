.PHONY: build run docker-build


build-nordigen-client:
	cd nordigenApiClient && \
	go build -o ./bin/go-nordigen-client && \
	chmod +x ./bin/go-nordigen-client 

run-nordigen-client: build-nordigen-client
	./nordigenApiClient/bin/go-nordigen-client

build-recorder:
	cd recorder && \
	go build -o ./bin/go-recorder && \
	chmod +x ./bin/go-recorder 

run-recorder: build-nordigen-client
	./recorder/bin/go-recorder

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
		-p 5432:5432 \
		-d postgres:16
stop-db:
	sudo docker stop my-db

remove-db:
	sudo docker rm my-db

start-rabitmq:
	sudo docker run -d --name rabbitmq -p 5672:5672 \ 
		-p 15672:15672 rabbitmq:3.12-management

stop-rabbitmq:
	sudo docker stop rabbitmq

remove-rabbitmq:
	sudo docker rm rabbitmq

start-grafana:
	sudo docker run -d --name=grafana -p 3000:3000 grafana/grafana

stop grafana:
	sudo docker stop grafana

remove-grafana:
	sudo docker rm grafana

