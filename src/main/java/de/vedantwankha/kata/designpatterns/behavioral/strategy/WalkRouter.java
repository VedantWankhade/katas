package de.vedantwankha.kata.designpatterns.behavioral.strategy;

public class WalkRouter implements Router {
    @Override
    public String route() {
        return "WALKING";
    }
}
