module github.com/q-sw/go-bank-analysis/nordigen-api-client

go 1.21.3

replace github.com/q-sw/go-bank-analysis/utils v0.0.0 => ../utils

replace github.com/q-sw/go-bank-analysis/types v0.0.0 => ../types

require (
	github.com/gorilla/mux v1.8.1
	github.com/q-sw/go-bank-analysis/types v0.0.0
	github.com/q-sw/go-bank-analysis/utils v0.0.0
)

require github.com/rabbitmq/amqp091-go v1.9.0
