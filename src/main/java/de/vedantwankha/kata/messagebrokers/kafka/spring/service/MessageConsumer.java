package de.vedantwankha.kata.messagebrokers.kafka.spring.service;

import de.vedantwankha.kata.messagebrokers.kafka.spring.Message;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Component
public class MessageConsumer {
    @Autowired
    private WebsocketService socket;

    @KafkaListener(topics = "test-topic", groupId = "test-topic")
    public void listen(Message msg) throws Exception {
        System.out.println("sending over socket: " + msg);
        socket.message(msg);
    }
}
