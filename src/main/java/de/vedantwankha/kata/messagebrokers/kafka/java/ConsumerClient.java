package de.vedantwankha.kata.messagebrokers.kafka.java;

import org.apache.kafka.clients.consumer.Consumer;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.KafkaConsumer;

import java.io.IOException;
import java.time.Duration;
import java.util.Collections;
import java.util.Properties;

public class ConsumerClient {
    public static void main(String[] args) {
        var config = new Properties();
        try(var in = ConsumerClient.class.getClassLoader().getResourceAsStream("config.properties")) {
            config.load(in);
            Consumer<String, String> consumer = new KafkaConsumer<>(config);
            consumer.subscribe(Collections.singleton(config.getProperty("kafka.topic")));
            while (true) {
                ConsumerRecords<String, String> records = consumer.poll(Duration.ofMillis(100));
                for (ConsumerRecord<String, String> record : records)
                    System.out.printf("offset = %d, key = %s, value = %s%n", record.offset(), record.key(), record.value());
            }
//            consumer.close();
        } catch (NullPointerException | IOException e) {
            System.out.println("File config.properties does not exist");
            System.out.println("Exiting...");
        }
    }
}
