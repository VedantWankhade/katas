package de.vedantwankha.kata.messagebrokers.kafka.spring.controller;

import de.vedantwankha.kata.messagebrokers.kafka.spring.Message;
import de.vedantwankha.kata.messagebrokers.kafka.spring.service.MessageConsumer;
import de.vedantwankha.kata.messagebrokers.kafka.spring.service.MessageProducer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;

@org.springframework.stereotype.Controller
public class Controller {

    @Autowired
    private MessageProducer producer;

    @Autowired
    private MessageConsumer consumer;

    List<String> list = new ArrayList<>();

    @GetMapping("/producer")
    public String producer(@RequestParam(name = "topic") String topic, Model model) {
        model.addAttribute("topic", topic);
        return "producer";
    }

    @PostMapping("/producer")
    public String producerPost(@RequestParam(name="topic") String topic, @ModelAttribute Message message, Model model) {
        model.addAttribute("topic", topic);
        producer.sendMessage(topic, message.message());
        return "producer";
    }

    @GetMapping("/consumer")
    public String consumer(@RequestParam(name = "topic") String topic, Model model) {
        model.addAttribute("topic", topic);
        return "consumer";
    }
}
