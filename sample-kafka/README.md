
# sample-kafka

This directory contains a minimal environment to quickly spin up a local Kafka cluster using Docker Compose, ideal for testing and experimentation.

## Structure

- **docker-compose.yaml**: Spins up the Kafka and Zookeeper environment.
- **commands.txt**: Useful commands for managing the environment and topics.

## Useful Commands

```sh
First, start the compose
    docker compose up

Enter the container terminal
    docker exec -it kafka_container_name bash

To create a topic
    kafka-topics --create --bootstrap-server=localhost:9092 --topic=topic_name --partitions=2 --replication-factor=1

To remove a topic
    kafka-topics --bootstrap-server=localhost:9092 --delete --topic=topic_name

To describe a topic
    kafka-topics --bootstrap-server=localhost:9092 --topic=topic_name --describe

To start a Consumer
    kafka-console-consumer --bootstrap-server=localhost:9092 --topic=topic_name --group=group_name
    you can pass a group to the consumer using --group=group_name

To start a Producer
    kafka-console-producer --bootstrap-server=localhost:9092 --topic=topic_name

To check Consumer-group
    kafka-consumer-groups --bootstrap-server=localhost:9092 --group=group_name --describe
```

---

> See the main monorepo README for more general information.
