package de.vedantwankha.kata.messagebrokers.kafka.spring.service;

import de.vedantwankha.kata.messagebrokers.kafka.spring.Message;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.messaging.simp.SimpMessagingTemplate;
import org.springframework.stereotype.Service;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

@Service
public class WebsocketService {
    private final SimpMessagingTemplate msgTmp;

    @Autowired
    public WebsocketService(SimpMessagingTemplate msgTmp) {
        this.msgTmp = msgTmp;
    }

    @RequestMapping(path="/messages", method= RequestMethod.POST)
    public void message(Message message) throws Exception {
        System.out.println(message.message());
        msgTmp.convertAndSend("/topic/messages", message.message());
    }
}
