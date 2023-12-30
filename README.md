# go-bank-analysis
GO API to analyse my bank accounts

## Instroduction

This project use Nordigen API to fectch my bank accounts opperation.
All of my operations will be store in a database ([PostgreSQL](https://www.postgresql.org/docs/)) and the result will be visualize with [Grafana](https://grafana.com/docs/).

All of the componment will be run in a kubernetes cluster. For my test I use [minikube](https://minikube.sigs.k8s.io/docs/)

## How to use this projet

To build and run this project you can use the `makefile` rules described below:

| Rule | Description |command|
|------|-------------|-------|
|build| build the Go binary|`make build`|
|run| Run localy the binary|`make run`|
|docker-build| build the docker image|`make docker-build`|

