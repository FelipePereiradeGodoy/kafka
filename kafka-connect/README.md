# kafka-connect

Kafka Connect study environment, including connector configuration examples for MongoDB and MySQL, as well as data persistence and usage instructions.

## Useful Links

- [Kafka Control Center](http://localhost:9021/)
- [Mongo Express](http://localhost:8085/)

## Notes

- **Kafka Confluent**: The hub for plugins to work with Kafka Connect (source, sink, transform, converter).

### DLQ - Dead Letter Queue

- `errors.tolerance = none`: makes the task fail immediately along with the error, i.e., it blocks the queue. **[DEFAULT]**
- `errors.tolerance = all`: errors are ignored and the process continues normally.
- `errors.deadletterqueue.topic.name = <topic-name>`: must be used together with `tolerance=all` and errors go to the DLQ, so the process does not block and you can identify the error.

---

> See the main monorepo README for more general information.