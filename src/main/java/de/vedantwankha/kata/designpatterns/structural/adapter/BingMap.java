package de.vedantwankha.kata.designpatterns.structural.adapter;

// BingMap doesnt implement StandardImageProvider
public class BingMap {
    public String image(int y, int x) {
        return String.format("Image (%s, %s) from Bing", x, y);
    }
}
