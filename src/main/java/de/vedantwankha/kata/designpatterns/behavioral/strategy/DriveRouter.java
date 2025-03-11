package de.vedantwankha.kata.designpatterns.behavioral.strategy;

public class DriveRouter implements Router {
    @Override
    public String route() {
        return "DRIVING";
    }
}
