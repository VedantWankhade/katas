package de.vedantwankha.kata.designpatterns.behavioral.obsever;

public interface EventSubject {
    void addSubscriber(EventSubscriber sub);
//    void removeSubscriber(EventSubscriber sub);
    void notifySubscribers();
}

