# Atlas

Middleware for Storing Data on Multiple Blockchains

<!-- ![Atlas Icon](docs/images/icon.png?raw=true "Icon retrive from https://www.vectorstock.com/royalty-free-vector/atlas-holding-up-world-mascot-vector-20996478") -->

## What it is?

A blockchain is a distributed database that maintains a continuously growing list of ordered records, called blocks. These blocks are linked using cryptography. Each block contains a cryptographic hash of the previous block, a timestamp, and transaction data.

Our systems tries to make it simple to create links cross-chains. In other words, it allows blockchains to speak to one another in a standardized way. Cross blockchain compatibility, allows different blockchains to communicate with one another without the help of intermediaries. Instead, people will be able to transact with users from other compatible blockchains. The entire process will take place without any downtime or expensive transaction fees.

## How it works?

![Atlas Diagram](docs/images/BlockchainMiddleware.png?raw=true "Icon retrive from https://www.vectorstock.com/royalty-free-vector/atlas-holding-up-world-mascot-vector-20996478")

- [Available API routes](docs)
- [UML Diagram](docs/diagrams/UML.png)

## Running The Project

- Clone this project locally
- Run `make setup` in your bash / command line
- Open another terminal and run `make db-up` in your bash / command line
- Run `make db-migration` in your bash / command line
- Run `run-local`

To see all the available commands run `make help` in your bash / command line

```!/bin/bash
---------------- HELP ---------------------
db-up  Start DB Strucutre
db-migration  Create DB Strucutre
setup   Install Go Modules
run-local  Start project locally
unit-test  Start unit test
```

## Dependencies

- [Golang](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)
