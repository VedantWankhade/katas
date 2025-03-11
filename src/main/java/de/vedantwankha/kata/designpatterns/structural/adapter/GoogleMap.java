package de.vedantwankha.kata.designpatterns.structural.adapter;

public class GoogleMap implements StandardImageProvider {
    @Override
    public String getImage(int x, int y) {
        return String.format("Image (%s, %s) from Google", x, y);
    }
}
