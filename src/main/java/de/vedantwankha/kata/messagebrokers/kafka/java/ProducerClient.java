package de.vedantwankha.kata.messagebrokers.kafka.java;

import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.Producer;
import org.apache.kafka.clients.producer.ProducerRecord;

import java.io.IOException;
import java.util.Properties;

public class ProducerClient {
    public static void main(String[] args) {
        var config = new Properties();
        try(var in = ProducerClient.class.getClassLoader().getResourceAsStream("config.properties")) {
            config.load(in);
            Producer<String, String> producer = new KafkaProducer<>(config);
            for (int i = 0; i < 10; i++) {
                producer.send(new ProducerRecord<>(config.getProperty("kafka.topic"), Integer.toString(i), Integer.toString(i)));
            }
            producer.close();
        } catch (NullPointerException | IOException e) {
            System.out.println("File config.properties does not exist");
            System.out.println("Exiting...");
        }
    }
}
