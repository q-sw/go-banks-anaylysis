module github.com/q-sw/go-bank-analysis/recorder

go 1.21.3

require (
	github.com/lib/pq v1.10.9
	github.com/q-sw/go-bank-analysis/types v0.0.0
	github.com/rabbitmq/amqp091-go v1.9.0
)

replace github.com/q-sw/go-bank-analysis/types v0.0.0 => ../types
