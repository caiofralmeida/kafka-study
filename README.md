# Kafka Study

### How to setup Kafka single broker stack

The command bellow will startup the Kafka stack (zookeeper, kafka server, kafka-manager and kafdrop) and it will create a new topic, you can change the topic name in `KAFKA_CREATE_TOPICS` var env
or creating a new one by using the kafka command cli.
```
docker-compose up -d
```

### Starting the consumer

Open a new terminal and run the command bellow to start the consumer.
```
go run consumer.go
```

### Producing a new message to the topic

Change the topic name into `producer.go` file if it's necessary. Open a new terminal and run the command bellow giving the key and value to the command.
```
go run producer.go <key> "<message>"
```
