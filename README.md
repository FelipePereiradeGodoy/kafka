
# Kafka Study Monorepo

This repository is a **monorepo** that brings together several projects and practical examples for studying and experimenting with Apache Kafka and its ecosystem. The goal is to centralize different approaches, languages, and tools related to Kafka, making learning and code reuse easier.

## Monorepo Structure

- **go-kafka/**: Producer, consumer, and recovery examples written in Go, with Docker and docker-compose for easy local execution.
- **kafka-connect/**: Study environment for Kafka Connect, including connector configuration examples for MongoDB and MySQL, as well as data persistence and usage instructions.
- **sample-kafka/**: Basic commands and examples to quickly spin up a local Kafka environment via docker-compose.

## How to Use

Each subdirectory has its own README or instruction files. It is recommended to navigate to the desired directory and follow the specific instructions. Below is a quick summary:

### go-kafka

- **cmd/producer/**: Kafka producer example in Go.
- **cmd/consumer/**: Kafka consumer example in Go.
- **cmd/recovery/**: Message recovery example.
- **docker-compose.yaml**: Spins up the Kafka environment for local testing.
- **commands.txt**: Useful commands for build, execution, and testing.

### kafka-connect

- **docker-compose.yaml**: Spins up the Kafka Connect environment, MongoDB, and other required services.
- **connectors/**: Connector configuration examples (MongoDB, MySQL).
- **README.md**: Detailed usage instructions.
- **data/**: Folder for sample data.

### sample-kafka

- **docker-compose.yaml**: Basic Kafka environment for quick tests.
- **commands.txt**: Useful commands for environment management.

## Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/) (for Go examples)

## Recommendations

- Always read the `commands.txt` of each subproject for useful commands and usage examples.
- Check the specific READMEs for configuration and execution details.
- Feel free to modify, create new examples, or suggest improvements!

---

> **This repository is intended for educational and experimentation purposes.**
