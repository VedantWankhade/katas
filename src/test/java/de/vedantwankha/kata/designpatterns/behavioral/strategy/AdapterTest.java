package de.vedantwankha.kata.designpatterns.behavioral.strategy;

import de.vedantwankha.kata.designpatterns.structural.adapter.BingMap;
import de.vedantwankha.kata.designpatterns.structural.adapter.BingToStandardAdapter;
import de.vedantwankha.kata.designpatterns.structural.adapter.GoogleMap;
import de.vedantwankha.kata.designpatterns.structural.adapter.Map;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

@DisplayName("Test Adapter pattern")
public class AdapterTest {
    @Test
    @DisplayName("Test with class that does not need adapter")
    void test1() {
        // no adapter needed
        Map m = new Map(new GoogleMap());
        m.showImage(1, 2);
    }

    @Test
    @DisplayName("Test with adaptee without the adapter")
    void test2() {
        // wont work
//        Map m = new Map(new BingMap());
//        m.showImage(1, 2);
    }

    @Test
    @DisplayName("Test adaptee with the adapter")
    void test3() {
        // BingMap doesnt implement StandardImageProvider but BingToStandardAdapter does
        Map m = new Map(new BingToStandardAdapter());
        m.showImage(1, 2);
    }
}
