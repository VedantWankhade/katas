package de.vedantwankha.kata.designpatterns.behavioral.strategy;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

@DisplayName("Test strategy pattern")
class MapTest {
    @Test
    @DisplayName("Provide implementation of Router while creating map object")
    void testStrategy1() {
        Map map = new Map(new DriveRouter());
        map.route("A", "B");
    }

    @Test
    @DisplayName("Provide implementation of Router at runtime after creating object")
    void testStrategy2() {
        Map map = new Map(new DriveRouter());
        map.route("A", "B");
        map.setRouter(new WalkRouter());
        map.route("A", "B");
    }

    @Test
    @DisplayName("Provide implementation of a custom Router")
    void testStrategy3() {
        Map map = new Map(new Router() {
            @Override
            public String route() {
                return "FLYING";
            }
        });
        map.route("A", "B");

        map.setRouter(() -> "BUNNY HOPPING");
        map.route("A", "B");
    }
}